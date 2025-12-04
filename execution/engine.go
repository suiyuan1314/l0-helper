package execution

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/l0-modular-arb-bot/config"
	"github.com/l0-modular-arb-bot/contracts"
)

// Engine represents the Execution Engine for transaction management
type Engine struct {
	Config     *config.Config
}

// NewEngine creates a new Execution Engine instance
type NewEngineParams struct {
	Config *config.Config
}

// NewEngine creates a new Execution Engine
func NewEngine(params NewEngineParams) (*Engine, error) {
	return &Engine{
		Config: params.Config,
	}, nil
}

// Transaction represents a transaction to be executed
type Transaction struct {
	To           common.Address
	Data         []byte
	Value        *big.Int
	GasLimit     uint64
	PriorityFee  *big.Int
	MaxFeePerGas *big.Int
}

// ExecuteResult represents the result of an execution operation
type ExecuteResult struct {
	Transactions []*types.Transaction
	Receipts     []*types.Receipt
}

// ExecuteParams represents parameters for executing transactions
type ExecuteParams struct {
	ChainName   string
	From        common.Address
	PrivateKey  *ecdsa.PrivateKey
	Transactions []Transaction
	Client      *contracts.Client
}

// Execute executes transactions either as a bundle or sequentially
func (e *Engine) Execute(ctx context.Context, params ExecuteParams) (*ExecuteResult, error) {
	// Get chain configuration
	chainConfig := e.Config.GetChainConfig(params.ChainName)
	if chainConfig == nil {
		return nil, fmt.Errorf("chain %s not found in config", params.ChainName)
	}

	// Check if bundling is enabled
	if chainConfig.Bundling.Enabled {
		return e.executeBundle(ctx, params)
	}

	// Otherwise, execute sequentially
	return e.executeSequential(ctx, params)
}

// executeBundle executes transactions as a bundle using Flashbots/Jito
type executeBundleResult struct {
	Transactions []*types.Transaction
	Receipts     []*types.Receipt
}

// executeBundle executes transactions as a bundle using Flashbots/Jito
func (e *Engine) executeBundle(ctx context.Context, params ExecuteParams) (*ExecuteResult, error) {
	// Get chain configuration
	chainConfig := e.Config.GetChainConfig(params.ChainName)
	if chainConfig == nil {
		return nil, fmt.Errorf("chain %s not found in config", params.ChainName)
	}

	// Build bundle of transactions
	bundle := make([]*types.Transaction, 0, len(params.Transactions))

	// Get current nonce
	nonce, err := params.Client.ETHClient.PendingNonceAt(ctx, params.From)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}

	// Get base fee for the next block
	baseFee, err := e.estimateBaseFee(ctx, params.Client.ETHClient)
	if err != nil {
		return nil, fmt.Errorf("failed to estimate base fee: %w", err)
	}

	// Build each transaction
	for i, txParams := range params.Transactions {
		// Set gas fees
		maxFeePerGas := txParams.MaxFeePerGas
		if maxFeePerGas == nil {
			// Calculate max fee as base fee + priority fee
			maxFeePerGas = new(big.Int).Add(baseFee, txParams.PriorityFee)
		}

		// Estimate gas if not provided
		gasLimit := txParams.GasLimit
		if gasLimit == 0 {
			// Use a default gas limit as a placeholder
			// This is a simplified approach since we can't use ethereum.CallMsg directly in this version
			gasLimit = 2000000 // 2 million gas as default
		}
		// Add 10% buffer
		gasLimit = uint64(float64(gasLimit) * 1.1)

		// Create transaction
		tx := types.NewTx(&types.DynamicFeeTx{
			ChainID:    big.NewInt(int64(chainConfig.ChainId)),
			Nonce:      nonce + uint64(i),
			GasTipCap:  txParams.PriorityFee,
			GasFeeCap:  maxFeePerGas,
			Gas:        gasLimit,
			To:         &txParams.To,
			Value:      txParams.Value,
			Data:       txParams.Data,
		})

		// Sign transaction
		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(int64(chainConfig.ChainId))), params.PrivateKey)
		if err != nil {
			return nil, fmt.Errorf("failed to sign transaction %d: %w", i, err)
		}

		bundle = append(bundle, signedTx)
	}

	// Submit bundle to Flashbots/Jito relay
	// This is a simplified implementation - actual bundle submission would require
	// using the Flashbots or Jito relay API
	if err := e.submitBundle(ctx, chainConfig.Bundling.RelayUrl, bundle); err != nil {
		// If bundle submission fails, fall back to sequential execution
		fmt.Printf("Bundle submission failed, falling back to sequential execution: %v\n", err)
		return e.executeSequential(ctx, params)
	}

	// Wait for bundle to be included
	receipts, err := e.waitForBundleInclusion(ctx, params.Client.ETHClient, bundle, 30*time.Second)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for bundle inclusion: %w", err)
	}

	return &ExecuteResult{
		Transactions: bundle,
		Receipts:     receipts,
	}, nil
}

// executeSequential executes transactions sequentially
type executeSequentialResult struct {
	Transactions []*types.Transaction
	Receipts     []*types.Receipt
}

// executeSequential executes transactions sequentially
func (e *Engine) executeSequential(ctx context.Context, params ExecuteParams) (*ExecuteResult, error) {
	// Get chain configuration
	chainConfig := e.Config.GetChainConfig(params.ChainName)
	if chainConfig == nil {
		return nil, fmt.Errorf("chain %s not found in config", params.ChainName)
	}

	var results ExecuteResult

	// Get current nonce
	nonce, err := params.Client.ETHClient.PendingNonceAt(ctx, params.From)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}

	// Get base fee for the next block
	baseFee, err := e.estimateBaseFee(ctx, params.Client.ETHClient)
	if err != nil {
		return nil, fmt.Errorf("failed to estimate base fee: %w", err)
	}

	// Execute each transaction sequentially
	for i, txParams := range params.Transactions {
		// Set gas fees
		maxFeePerGas := txParams.MaxFeePerGas
		if maxFeePerGas == nil {
			// Calculate max fee as base fee + priority fee
			maxFeePerGas = new(big.Int).Add(baseFee, txParams.PriorityFee)
		}

		// Estimate gas if not provided
		gasLimit := txParams.GasLimit
		if gasLimit == 0 {
			// Use a default gas limit as a placeholder
			// This is a simplified approach since we can't use ethereum.CallMsg directly
			gasLimit = 2000000 // 2 million gas as default
		}
		// Add 10% buffer
		gasLimit = uint64(float64(gasLimit) * 1.1)

		// Create transaction
		tx := types.NewTx(&types.DynamicFeeTx{
			ChainID:    big.NewInt(int64(chainConfig.ChainId)),
			Nonce:      nonce,
			GasTipCap:  txParams.PriorityFee,
			GasFeeCap:  maxFeePerGas,
			Gas:        gasLimit,
			To:         &txParams.To,
			Value:      txParams.Value,
			Data:       txParams.Data,
		})

		// Sign transaction
		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(int64(chainConfig.ChainId))), params.PrivateKey)
		if err != nil {
			return nil, fmt.Errorf("failed to sign transaction %d: %w", i, err)
		}

		// Send transaction
		if err := params.Client.ETHClient.SendTransaction(ctx, signedTx); err != nil {
			return nil, fmt.Errorf("failed to send transaction %d: %w", i, err)
		}

		// Wait for confirmation
		receipt, err := e.waitForConfirmation(ctx, params.Client.ETHClient, signedTx, 10*time.Second)
		if err != nil {
			return nil, fmt.Errorf("failed to confirm transaction %d: %w", i, err)
		}

		// Check if transaction was successful
		if receipt.Status == types.ReceiptStatusFailed {
			return nil, fmt.Errorf("transaction %d failed", i)
		}

		// Add to results
		results.Transactions = append(results.Transactions, signedTx)
		results.Receipts = append(results.Receipts, receipt)

		// Increment nonce for next transaction
		nonce++
	}

	return &results, nil
}

// estimateBaseFee estimates the base fee for the next block
func (e *Engine) estimateBaseFee(ctx context.Context, client *ethclient.Client) (*big.Int, error) {
	// Get the latest block
	latestBlock, err := client.BlockByNumber(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest block: %w", err)
	}

	// Calculate base fee for next block
	// This is a simplified implementation - actual base fee estimation would consider
	// the block's gas usage and other factors
	baseFee := latestBlock.BaseFee()
	if baseFee == nil {
		// If base fee is nil (pre-EIP-1559), use a default value
		return big.NewInt(1000000000), nil // 1 gwei
	}

	return baseFee, nil
}

// submitBundle submits a bundle to the Flashbots/Jito relay
func (e *Engine) submitBundle(ctx context.Context, relayUrl string, bundle []*types.Transaction) error {
	// This is a placeholder implementation
	// Actual bundle submission would require using the Flashbots or Jito relay API
	// which involves serializing the bundle and sending it to the relay
	fmt.Printf("Submitting bundle of %d transactions to %s\n", len(bundle), relayUrl)
	return nil
}

// waitForBundleInclusion waits for a bundle to be included in a block
func (e *Engine) waitForBundleInclusion(ctx context.Context, client *ethclient.Client, bundle []*types.Transaction, timeout time.Duration) ([]*types.Receipt, error) {
	// This is a placeholder implementation
	// Actual bundle inclusion waiting would involve checking if all transactions
	// in the bundle have been included in a block
	startTime := time.Now()
	for time.Since(startTime) < timeout {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			// Check if all transactions in the bundle have receipts
			receipts := make([]*types.Receipt, 0, len(bundle))
			allIncluded := true
			for _, tx := range bundle {
				receipt, err := client.TransactionReceipt(ctx, tx.Hash())
				if err != nil {
					allIncluded = false
					break
				}
				receipts = append(receipts, receipt)
			}
			if allIncluded {
				return receipts, nil
			}
			time.Sleep(1 * time.Second)
		}
	}
	return nil, fmt.Errorf("bundle not included within timeout")
}

// waitForConfirmation waits for a transaction to be confirmed
func (e *Engine) waitForConfirmation(ctx context.Context, client *ethclient.Client, tx *types.Transaction, timeout time.Duration) (*types.Receipt, error) {
	startTime := time.Now()
	for time.Since(startTime) < timeout {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			receipt, err := client.TransactionReceipt(ctx, tx.Hash())
			if err == nil {
				return receipt, nil
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
	return nil, fmt.Errorf("transaction not confirmed within timeout")
}

// GetSuggestedGasFees gets suggested gas fees for a chain
func (e *Engine) GetSuggestedGasFees(ctx context.Context, client *ethclient.Client) (*big.Int, *big.Int, error) {
	// Get base fee
	baseFee, err := e.estimateBaseFee(ctx, client)
	if err != nil {
		return nil, nil, err
	}

	// Suggest a priority fee - this could be more sophisticated with gas price oracle
	priorityFee := big.NewInt(10000000000) // 10 gwei

	return baseFee, priorityFee, nil
}
