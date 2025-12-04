package guard

import (
	"context"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/l0-modular-arb-bot/config"
	"github.com/l0-modular-arb-bot/contracts"
)

// Engine represents the Guard Engine for risk validation
type Engine struct {
	Config     *config.Config
}

// NewEngine creates a new Guard Engine instance
type NewEngineParams struct {
	Config *config.Config
}

// NewEngine creates a new Guard Engine
func NewEngine(params NewEngineParams) (*Engine, error) {
	return &Engine{
		Config: params.Config,
	}, nil
}

// CheckExecutabilityParams represents parameters for checking executability
type CheckExecutabilityParams struct {
	ChainName   string
	Origin      Origin
	Receiver    common.Address
	Client      *contracts.Client
}

// Origin represents the origin of a LayerZero message
type Origin struct {
	SrcEid  uint32
	Sender  common.Hash
	Nonce   uint64
}

// ExecutabilityStatus represents the executability status of a message
type ExecutabilityStatus int

const (
	ExecutabilityStatusNotExecutable ExecutabilityStatus = iota
	ExecutabilityStatusExecutable
	ExecutabilityStatusExecuted
)

// CheckExecutability checks if a message is executable
type CheckExecutabilityResult struct {
	Status ExecutabilityStatus
}

// CheckExecutability checks if a message is executable
func (e *Engine) CheckExecutability(ctx context.Context, params CheckExecutabilityParams) (*CheckExecutabilityResult, error) {
	// Get chain configuration
	chainConfig := e.Config.GetChainConfig(params.ChainName)
	if chainConfig == nil {
		return nil, fmt.Errorf("chain %s not found in config", params.ChainName)
	}

	// Check if EndpointView is configured
	if chainConfig.Contracts.EndpointView == "" {
		return nil, fmt.Errorf("EndpointView contract not configured for chain %s", params.ChainName)
	}

	// Call EndpointV2View.executable()
	args := []interface{}{
		params.Origin.SrcEid,
		params.Origin.Sender,
		params.Origin.Nonce,
		params.Receiver,
	}

	result, err := params.Client.CallContract(ctx, common.HexToAddress(chainConfig.Contracts.EndpointView), params.Client.EndpointV2ViewABI, "executable", args...)
	if err != nil {
		return nil, fmt.Errorf("failed to call executable: %w", err)
	}

	// Parse result
	if len(result) != 1 {
		return nil, fmt.Errorf("invalid executable result length: %d", len(result))
	}

	status, ok := result[0].(uint8)
	if !ok {
		return nil, fmt.Errorf("failed to parse executability status")
	}

	return &CheckExecutabilityResult{
		Status: ExecutabilityStatus(status),
	}, nil
}

// CheckNonceGapParams represents parameters for checking nonce gaps
type CheckNonceGapParams struct {
	ChainName   string
	Origin      Origin
	Receiver    common.Address
	Client      *contracts.Client
}

// CheckNonceGapResult represents the result of a nonce gap check
type CheckNonceGapResult struct {
	Gap          uint64
	Cost         *big.Int
	CanUnblock   bool
}

// CheckNonceGap checks for nonce gaps and calculates the cost to unblock
func (e *Engine) CheckNonceGap(ctx context.Context, params CheckNonceGapParams) (*CheckNonceGapResult, error) {
	// Get chain configuration
	chainConfig := e.Config.GetChainConfig(params.ChainName)
	if chainConfig == nil {
		return nil, fmt.Errorf("chain %s not found in config", params.ChainName)
	}

	// Get current inbound nonce
	args := []interface{}{
		params.Receiver,
		params.Origin.SrcEid,
		params.Origin.Sender,
	}

	result, err := params.Client.CallContract(ctx, common.HexToAddress(chainConfig.Contracts.Endpoint), params.Client.EndpointV2ABI, "inboundNonce", args...)
	if err != nil {
		return nil, fmt.Errorf("failed to call inboundNonce: %w", err)
	}

	// Parse result
	if len(result) != 1 {
		return nil, fmt.Errorf("invalid inboundNonce result length: %d", len(result))
	}

	currentNonce, ok := result[0].(uint64)
	if !ok {
		return nil, fmt.Errorf("failed to parse inboundNonce")
	}

	// Calculate gap
	gap := params.Origin.Nonce - currentNonce

	// Check if gap is within threshold
	canUnblock := gap <= e.Config.Bot.MaxGapThreshold && gap > 0

	// Calculate cost to unblock (simplified implementation)
	// In a real implementation, this would estimate the gas cost for each unblocking transaction
	cost := big.NewInt(0)
	if canUnblock {
		// Estimate gas cost per lzReceive transaction
		gasPerTx := big.NewInt(500000) // 500k gas per tx
		gasPrice := big.NewInt(10000000000) // 10 gwei
		costPerTx := new(big.Int).Mul(gasPerTx, gasPrice)
		cost = new(big.Int).Mul(costPerTx, big.NewInt(int64(gap)))
	}

	return &CheckNonceGapResult{
		Gap:        gap,
		Cost:       cost,
		CanUnblock: canUnblock,
	}, nil
}

// CheckProfitParams represents parameters for checking profit
type CheckProfitParams struct {
	ExpectedOut   *big.Int
	GasCost       *big.Int
	BridgeCost    *big.Int
	MinProfit     *big.Int
}

// CheckProfitResult represents the result of a profit check
type CheckProfitResult struct {
	IsProfitable bool
	Profit       *big.Int
}

// CheckProfit checks if a transaction is profitable
func (e *Engine) CheckProfit(params CheckProfitParams) (*CheckProfitResult, error) {
	// Calculate total cost
	totalCost := new(big.Int).Add(params.GasCost, params.BridgeCost)

	// Calculate profit
	profit := new(big.Int).Sub(params.ExpectedOut, totalCost)

	// Check if profit meets minimum threshold
	isProfitable := profit.Cmp(params.MinProfit) >= 0

	return &CheckProfitResult{
		IsProfitable: isProfitable,
		Profit:       profit,
	}, nil
}

// SimulateProfitParams represents parameters for simulating profit
type SimulateProfitParams struct {
	ChainName   string
	From        common.Address
	Transactions []Transaction
	Client      *contracts.Client
}

// Transaction represents a transaction to be simulated
type Transaction struct {
	To           common.Address
	Data         []byte
	Value        *big.Int
}

// SimulateProfitResult represents the result of a profit simulation
type SimulateProfitResult struct {
	IsProfitable bool
	Profit       *big.Int
	GasCost      *big.Int
}

// SimulateProfit simulates a transaction to verify profit
func (e *Engine) SimulateProfit(ctx context.Context, params SimulateProfitParams) (*SimulateProfitResult, error) {
	// Get chain configuration
	chainConfig := e.Config.GetChainConfig(params.ChainName)
	if chainConfig == nil {
		return nil, fmt.Errorf("chain %s not found in config", params.ChainName)
	}

	// This is a placeholder implementation
	// Actual profit simulation would involve:
	// 1. Forking the chain
	// 2. Simulating the transactions
	// 3. Calculating the profit based on state changes

	// For now, we'll use a simplified approach
	gasCost := big.NewInt(1000000000000000000) // 1 ETH gas cost
	profit := big.NewInt(50000000000000000) // 0.05 ETH profit
	isProfitable := profit.Cmp(big.NewInt(0)) > 0

	return &SimulateProfitResult{
		IsProfitable: isProfitable,
		Profit:       profit,
		GasCost:      gasCost,
	}, nil
}

// ValidateTransactionParams represents parameters for validating a transaction
type ValidateTransactionParams struct {
	ChainName   string
	From        common.Address
	To          common.Address
	Data        []byte
	Value       *big.Int
	Client      *contracts.Client
}

// ValidateTransactionResult represents the result of a transaction validation
type ValidateTransactionResult struct {
	IsValid bool
	Reason  string
}

// ValidateTransaction validates a transaction before execution
func (e *Engine) ValidateTransaction(ctx context.Context, params ValidateTransactionParams) (*ValidateTransactionResult, error) {
	// Estimate gas using direct RPC call
	gasEstimateArgs := map[string]interface{}{
		"from": params.From,
		"to":   params.To,
		"data": params.Data,
		"value": params.Value,
	}

	// Estimate gas with direct RPC call
	var gasLimit string
	err := params.Client.ETHClient.Client().CallContext(ctx, &gasLimit, "eth_estimateGas", gasEstimateArgs)
	if err != nil {
		return &ValidateTransactionResult{
			IsValid: false,
			Reason:  fmt.Sprintf("gas estimation failed: %v", err),
		}, nil
	}

	// Convert gas limit from hex string to uint64
	gasLimitInt, err := strconv.ParseUint(gasLimit[2:], 16, 64) // Remove 0x prefix
	if err != nil {
		return &ValidateTransactionResult{
			IsValid: false,
			Reason:  fmt.Sprintf("failed to parse gas limit: %v", err),
		}, nil
	}

	// Check if gas limit is reasonable
	if gasLimitInt > 10000000 { // 10 million gas limit
		return &ValidateTransactionResult{
			IsValid: false,
			Reason:  fmt.Sprintf("gas limit too high: %d", gasLimitInt),
		}, nil
	}

	// Check if transaction would fail (simulate it with eth_call)
	callArgs := map[string]interface{}{
		"from": params.From,
		"to":   params.To,
		"data": params.Data,
		"value": params.Value,
	}

	var result string
	err = params.Client.ETHClient.Client().CallContext(ctx, &result, "eth_call", callArgs, "latest")
	if err != nil {
		return &ValidateTransactionResult{
			IsValid: false,
			Reason:  fmt.Sprintf("transaction simulation failed: %v", err),
		}, nil
	}

	return &ValidateTransactionResult{
		IsValid: true,
		Reason:  "transaction is valid",
	}, nil
}
