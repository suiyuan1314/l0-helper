package strategy

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/l0-modular-arb-bot/bridge"
	"github.com/l0-modular-arb-bot/config"
	"github.com/l0-modular-arb-bot/contracts"
	"github.com/l0-modular-arb-bot/contracts/endpointv2"
	"github.com/l0-modular-arb-bot/contracts/endpointv2view"
	"github.com/l0-modular-arb-bot/execution"
	"github.com/l0-modular-arb-bot/guard"
	"github.com/l0-modular-arb-bot/liquidity"
)

// Controller represents the Strategy Controller for recipe composition
type Controller struct {
	Config          *config.Config
	LiquidityEngine *liquidity.Engine
	BridgeEngine    *bridge.Engine
	ExecutionEngine *execution.Engine
	GuardEngine     *guard.Engine
}

// NewController creates a new Strategy Controller instance
type NewControllerParams struct {
	Config          *config.Config
	LiquidityEngine *liquidity.Engine
	BridgeEngine    *bridge.Engine
	ExecutionEngine *execution.Engine
	GuardEngine     *guard.Engine
}

// NewController creates a new Strategy Controller
func NewController(params NewControllerParams) (*Controller, error) {
	return &Controller{
		Config:          params.Config,
		LiquidityEngine: params.LiquidityEngine,
		BridgeEngine:    params.BridgeEngine,
		ExecutionEngine: params.ExecutionEngine,
		GuardEngine:     params.GuardEngine,
	}, nil
}

// RecipeType represents the type of arbitrage recipe
type RecipeType int

const (
	RecipeTypeStandardArb RecipeType = iota
	RecipeTypeInventoryArb
	RecipeTypeRebalancing
	RecipeTypeTheCloser
	RecipeTypeTheCourier
	RecipeTypePanicSell
)

// Recipe represents an arbitrage recipe
type Recipe struct {
	Type RecipeType
}

// ExecuteRecipeParams represents parameters for executing a recipe
type ExecuteRecipeParams struct {
	Recipe           Recipe
	ChainName        string // Source chain for Standard/Inventory/Rebalancing, Dest for Closer/Panic
	From             common.Address
	PrivateKey       *ecdsa.PrivateKey
	Amount           *big.Int
	Client           *contracts.Client
	DestinationChain string
	DestEID          uint32
	ToAddress        common.Address
}

// ExecuteRecipeResult represents the result of executing a recipe
type ExecuteRecipeResult struct {
	Success      bool
	Message      string
	Transactions []*types.Transaction
}

// ExecuteRecipe executes an arbitrage recipe
func (c *Controller) ExecuteRecipe(ctx context.Context, params ExecuteRecipeParams) (*ExecuteRecipeResult, error) {
	switch params.Recipe.Type {
	case RecipeTypeStandardArb:
		return c.executeStandardArb(ctx, params)
	case RecipeTypeInventoryArb:
		return c.executeInventoryArb(ctx, params)
	case RecipeTypeRebalancing:
		return c.executeRebalancing(ctx, params)
	case RecipeTypeTheCloser:
		return c.executeTheCloser(ctx, params)
	case RecipeTypeTheCourier:
		return c.executeTheCourier(ctx, params)
	case RecipeTypePanicSell:
		return c.executePanicSell(ctx, params)
	default:
		return nil, fmt.Errorf("invalid recipe type: %d", params.Recipe.Type)
	}
}

// executeStandardArb executes the Standard Arb recipe: Buy → Bridge → Sell
func (c *Controller) executeStandardArb(ctx context.Context, params ExecuteRecipeParams) (*ExecuteRecipeResult, error) {
	// Step 1: Get swap quote from Liquidity Engine
	// Get OFT address from configuration
	chainConfig := c.Config.GetChainConfig(params.ChainName)
	if chainConfig == nil {
		return nil, fmt.Errorf("chain %s not found in config", params.ChainName)
	}
	swapQuote, err := c.LiquidityEngine.GetSwapQuote(ctx, liquidity.GetSwapQuoteParams{
		ChainName: params.ChainName,
		TokenIn:   common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"), // ETH
		TokenOut:  common.HexToAddress(chainConfig.Contracts.OFT),
		AmountIn:  params.Amount,
		Slippage:  0.005, // 0.5%
		Sender:    params.From,
		Recipient: params.From,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get swap quote: %w", err)
	}

	// Step 2: Send LayerZero message via Bridge Engine
	sendResult, err := c.BridgeEngine.Send(ctx, params.Client, bridge.SendParams{
		ChainName: params.ChainName,
		From:      params.From,
		ToAddress: params.ToAddress,
		Amount:    swapQuote.ExpectedOut,
		DstEID:    params.DestEID,
		FeeBuffer: c.Config.Bot.FeeBufferMultiplier,
	}, params.PrivateKey)
	if err != nil {
		// If bridge fails, trigger Panic Sell
		panicSellResult, panicSellErr := c.executePanicSell(ctx, params)
		if panicSellErr != nil {
			return &ExecuteRecipeResult{
				Success: false,
				Message: fmt.Sprintf("Standard Arb failed, Panic Sell also failed: %v", panicSellErr),
			}, nil
		}
		return panicSellResult, nil
	}

	// Step 3: Monitor for PacketVerified event on destination chain
	// This would typically be handled by a separate listener goroutine
	// For simplicity, we'll return after sending the bridge message
	return &ExecuteRecipeResult{
		Success:      true,
		Message:      "Standard Arb executed successfully",
		Transactions: []*types.Transaction{sendResult.Transaction},
	}, nil
}

// executeInventoryArb executes the Inventory Arb recipe: Bridge → Sell
func (c *Controller) executeInventoryArb(ctx context.Context, params ExecuteRecipeParams) (*ExecuteRecipeResult, error) {
	// Step 1: Send LayerZero message via Bridge Engine
	sendResult, err := c.BridgeEngine.Send(ctx, params.Client, bridge.SendParams{
		ChainName: params.ChainName,
		From:      params.From,
		ToAddress: params.ToAddress,
		Amount:    params.Amount,
		DstEID:    params.DestEID,
		FeeBuffer: c.Config.Bot.FeeBufferMultiplier,
	}, params.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to send bridge message: %w", err)
	}

	return &ExecuteRecipeResult{
		Success:      true,
		Message:      "Inventory Arb executed successfully",
		Transactions: []*types.Transaction{sendResult.Transaction},
	}, nil
}

// executeRebalancing executes the Rebalancing recipe: Bridge Only
func (c *Controller) executeRebalancing(ctx context.Context, params ExecuteRecipeParams) (*ExecuteRecipeResult, error) {
	// Step 1: Send LayerZero message via Bridge Engine
	sendResult, err := c.BridgeEngine.Send(ctx, params.Client, bridge.SendParams{
		ChainName: params.ChainName,
		From:      params.From,
		ToAddress: params.ToAddress,
		Amount:    params.Amount,
		DstEID:    params.DestEID,
		FeeBuffer: c.Config.Bot.FeeBufferMultiplier,
	}, params.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to send bridge message: %w", err)
	}

	return &ExecuteRecipeResult{
		Success:      true,
		Message:      "Rebalancing executed successfully",
		Transactions: []*types.Transaction{sendResult.Transaction},
	}, nil
}

// executeTheCloser executes The Closer recipe: Execute & Swap
type TheCloserParams struct {
	ChainName  string
	From       common.Address
	PrivateKey *ecdsa.PrivateKey
	Origin     endpointv2.Origin
	Client     *contracts.Client
	TokenIn    common.Address
	TokenOut   common.Address
}

// executeTheCloser executes The Closer recipe: Execute & Swap
func (c *Controller) executeTheCloser(ctx context.Context, params ExecuteRecipeParams) (*ExecuteRecipeResult, error) {
	// Step 1: Check if message is executable
	// NOTE: CheckExecutability requires endpointv2view.Origin (View contract struct)
	// We need to convert or construct correctly.

	executabilityResult, err := c.GuardEngine.CheckExecutability(ctx, guard.CheckExecutabilityParams{
		ChainName: params.ChainName,
		Origin: endpointv2view.Origin{
			SrcEid: params.DestEID,
			Sender: [32]byte{}, // Placeholder
			Nonce:  0,
		},
		Receiver: params.From,
		Client:   params.Client,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to check executability: %w", err)
	}

	if executabilityResult.Status != guard.ExecutabilityStatusExecutable {
		return &ExecuteRecipeResult{
			Success: false,
			Message: fmt.Sprintf("Message not executable: %d", executabilityResult.Status),
		}, nil
	}

	// Step 2: Check for nonce gaps
	nonceGapResult, err := c.GuardEngine.CheckNonceGap(ctx, guard.CheckNonceGapParams{
		ChainName: params.ChainName,
		Origin: endpointv2.Origin{
			SrcEid: params.DestEID,
			Sender: [32]byte{},
			Nonce:  0,
		},
		Receiver: params.From,
		Client:   params.Client,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to check nonce gap: %w", err)
	}

	// If there's a nonce gap, execute The Courier recipe first
	if nonceGapResult.CanUnblock {
		courierResult, err := c.executeTheCourier(ctx, params)
		if err != nil {
			return nil, fmt.Errorf("failed to execute The Courier: %w", err)
		}
		if !courierResult.Success {
			return courierResult, nil
		}
	}

	// Step 3: Get swap quote for destination chain
	destChainConfig := c.Config.GetChainConfig(params.ChainName)
	if destChainConfig == nil {
		return nil, fmt.Errorf("destination chain %s not found in config", params.ChainName)
	}
	swapQuote, err := c.LiquidityEngine.GetSwapQuote(ctx, liquidity.GetSwapQuoteParams{
		ChainName: params.ChainName,
		TokenIn:   common.HexToAddress(destChainConfig.Contracts.OFT),
		TokenOut:  common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"), // ETH
		AmountIn:  params.Amount,
		Slippage:  0.005, // 0.5%
		Sender:    params.From,
		Recipient: params.From,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get swap quote: %w", err)
	}

	// Step 4: Check Approval
	approveTx, err := c.ensureApproval(ctx, params.ChainName, common.HexToAddress(destChainConfig.Contracts.OFT), swapQuote.ToAddress, params.Amount, params.From, params.Client.ETHClient)
	if err != nil {
		return nil, fmt.Errorf("failed to ensure approval: %w", err)
	}

	// Step 5: Build transactions
	var transactions []execution.Transaction

	// LZ Receive
	transactions = append(transactions, execution.Transaction{
		To:    common.HexToAddress(c.Config.GetChainConfig(params.ChainName).Contracts.Endpoint),
		Data:  []byte{},
		Value: big.NewInt(0),
	})

	// Add approve if needed
	if approveTx != nil {
		transactions = append(transactions, *approveTx)
	}

	// Swap transaction
	transactions = append(transactions, execution.Transaction{
		To:    swapQuote.ToAddress,
		Data:  swapQuote.CallData,
		Value: big.NewInt(0),
	})

	// Step 6: Execute transactions
	executeResult, err := c.ExecutionEngine.Execute(ctx, execution.ExecuteParams{
		ChainName:    params.ChainName,
		From:         params.From,
		PrivateKey:   params.PrivateKey,
		Transactions: transactions,
		Client:       params.Client,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to execute transactions: %w", err)
	}

	return &ExecuteRecipeResult{
		Success:      true,
		Message:      "The Closer executed successfully",
		Transactions: executeResult.Transactions,
	}, nil
}

// executeTheCourier executes The Courier recipe: Pure Execute
func (c *Controller) executeTheCourier(ctx context.Context, params ExecuteRecipeParams) (*ExecuteRecipeResult, error) {
	// Step 1: Build lzReceive transaction
	transactions := []execution.Transaction{
		{
			To:    common.HexToAddress(c.Config.GetChainConfig(params.ChainName).Contracts.Endpoint),
			Data:  []byte{},
			Value: big.NewInt(0),
		},
	}

	// Step 2: Execute transaction
	executeResult, err := c.ExecutionEngine.Execute(ctx, execution.ExecuteParams{
		ChainName:    params.ChainName,
		From:         params.From,
		PrivateKey:   params.PrivateKey,
		Transactions: transactions,
		Client:       params.Client,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to execute transaction: %w", err)
	}

	return &ExecuteRecipeResult{
		Success:      true,
		Message:      "The Courier executed successfully",
		Transactions: executeResult.Transactions,
	}, nil
}

// executePanicSell executes The Panic Sell recipe: Source Cleanup
func (c *Controller) executePanicSell(ctx context.Context, params ExecuteRecipeParams) (*ExecuteRecipeResult, error) {
	// Step 1: Get swap quote to sell OFT back to ETH
	// Get OFT address from configuration
	chainConfig := c.Config.GetChainConfig(params.ChainName)
	if chainConfig == nil {
		return nil, fmt.Errorf("chain %s not found in config", params.ChainName)
	}
	swapQuote, err := c.LiquidityEngine.GetSwapQuote(ctx, liquidity.GetSwapQuoteParams{
		ChainName: params.ChainName,
		TokenIn:   common.HexToAddress(chainConfig.Contracts.OFT),
		TokenOut:  common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"), // ETH
		AmountIn:  params.Amount,
		Slippage:  0.01, // 1% (more aggressive for panic sell)
		Sender:    params.From,
		Recipient: params.From,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get swap quote: %w", err)
	}

	// Step 2: Ensure Approval
	approveTx, err := c.ensureApproval(ctx, params.ChainName, common.HexToAddress(chainConfig.Contracts.OFT), swapQuote.ToAddress, params.Amount, params.From, params.Client.ETHClient)
	if err != nil {
		return nil, fmt.Errorf("failed to ensure approval: %w", err)
	}

	// Step 3: Build swap transaction
	var transactions []execution.Transaction
	if approveTx != nil {
		transactions = append(transactions, *approveTx)
	}

	tx := execution.Transaction{
		To:    swapQuote.ToAddress,
		Data:  swapQuote.CallData,
		Value: big.NewInt(0),
	}
	transactions = append(transactions, tx)

	// Step 4: Execute transaction with high priority
	baseFee, priorityFee, err := c.ExecutionEngine.GetSuggestedGasFees(ctx, params.Client.ETHClient)
	if err != nil {
		return nil, fmt.Errorf("failed to get gas fees: %w", err)
	}

	// Increase priority fee by 5x for panic sell
	highPriorityFee := new(big.Int).Mul(priorityFee, big.NewInt(5))

	// Update all transactions gas fees
	for i := range transactions {
		transactions[i].PriorityFee = highPriorityFee
		transactions[i].MaxFeePerGas = new(big.Int).Add(baseFee, highPriorityFee)
	}

	// Step 5: Execute transaction
	executeResult, err := c.ExecutionEngine.Execute(ctx, execution.ExecuteParams{
		ChainName:    params.ChainName,
		From:         params.From,
		PrivateKey:   params.PrivateKey,
		Transactions: transactions,
		Client:       params.Client,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to execute panic sell: %w", err)
	}

	return &ExecuteRecipeResult{
		Success:      true,
		Message:      "Panic Sell executed successfully",
		Transactions: executeResult.Transactions,
	}, nil
}

// MonitorPackets monitors PacketVerified events and triggers appropriate recipes
type MonitorPacketsParams struct {
	ChainName string
	Client    *contracts.Client
	Handler   func(recipe Recipe, event bridge.PacketVerifiedEvent)
}

// MonitorPackets monitors PacketVerified events and triggers appropriate recipes
func (c *Controller) MonitorPackets(ctx context.Context, params MonitorPacketsParams) error {
	// Delegate to Bridge Engine's MonitorPacketVerified method
	return c.BridgeEngine.MonitorPacketVerified(ctx, bridge.MonitorPacketVerifiedParams{
		ChainName: params.ChainName,
		Client:    params.Client.ETHClient,
		Handler: func(event bridge.PacketVerifiedEvent) {
			// Determine which recipe to execute based on event data
			var recipe Recipe

			// For example, if the bot is the receiver, execute The Closer
			// This would be more sophisticated in a real implementation
			recipe.Type = RecipeTypeTheCloser

			// Call the handler with the determined recipe
			params.Handler(recipe, event)
		},
	})
}

// ensureApproval checks and generates approval transaction if needed
func (c *Controller) ensureApproval(ctx context.Context, chainName string, token common.Address, spender common.Address, amount *big.Int, owner common.Address, client *ethclient.Client) (*execution.Transaction, error) {
	// Check approval
	approved, err := c.LiquidityEngine.CheckApproval(ctx, liquidity.CheckApprovalParams{
		ChainName:      chainName,
		TokenAddress:   token,
		SpenderAddress: spender,
		OwnerAddress:   owner,
		RequiredAmount: amount,
		Client:         client,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to check approval: %w", err)
	}

	if approved {
		return nil, nil // No approval needed
	}

	// Generate approval
	data, to, err := c.LiquidityEngine.Approve(ctx, liquidity.ApproveParams{
		ChainName:      chainName,
		TokenAddress:   token,
		SpenderAddress: spender,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to generate approval: %w", err)
	}

	return &execution.Transaction{
		To:    to,
		Data:  data,
		Value: big.NewInt(0),
	}, nil
}
