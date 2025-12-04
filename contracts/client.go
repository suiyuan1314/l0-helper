package contracts

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Client represents a contract client for interacting with LayerZero V2 contracts
type Client struct {
	ETHClient      *ethclient.Client
	EndpointV2ABI  abi.ABI
	EndpointV2ViewABI abi.ABI
	OFTABI         abi.ABI
	ERC20ABI       abi.ABI
}

// NewClient creates a new contract client
type NewClientParams struct {
	RPCEndpoint string
}

// NewClient creates a new contract client instance
func NewClient(params NewClientParams) (*Client, error) {
	// Connect to Ethereum client
	client, err := ethclient.Dial(params.RPCEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RPC: %w", err)
	}

	// Load ABIs from files
	endpointV2ABI, err := loadABI("abis/EndpointV2.json")
	if err != nil {
		return nil, err
	}

	endpointV2ViewABI, err := loadABI("abis/EndpointV2View.json")
	if err != nil {
		return nil, err
	}

	oftABI, err := loadABI("abis/ioft.json")
	if err != nil {
		return nil, err
	}

	erc20ABI, err := loadABI("abis/ERC20.json")
	if err != nil {
		return nil, err
	}

	return &Client{
		ETHClient:      client,
		EndpointV2ABI:  endpointV2ABI,
		EndpointV2ViewABI: endpointV2ViewABI,
		OFTABI:         oftABI,
		ERC20ABI:       erc20ABI,
	}, nil
}

// loadABI loads an ABI from a file
func loadABI(filePath string) (abi.ABI, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return abi.ABI{}, fmt.Errorf("failed to open ABI file: %w", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return abi.ABI{}, fmt.Errorf("failed to read ABI file: %w", err)
	}

	var contractABI abi.ABI
	if err := json.Unmarshal(content, &contractABI); err != nil {
		return abi.ABI{}, fmt.Errorf("failed to unmarshal ABI: %w", err)
	}

	return contractABI, nil
}

// CallContract calls a contract method without modifying state
func (c *Client) CallContract(ctx context.Context, contractAddr common.Address, contractABI abi.ABI, method string, args ...interface{}) ([]interface{}, error) {
	data, err := contractABI.Pack(method, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to pack method %s: %w", method, err)
	}

	// Create a call message with just To and Data fields
	// Using a simple struct with the fields that CallContract expects
	callArgs := map[string]interface{}{
		"to":   contractAddr,
		"data": data,
	}

	// Use direct RPC call to eth_call
	var result []byte
	err = c.ETHClient.Client().CallContext(ctx, &result, "eth_call", callArgs, "latest")
	if err != nil {
		return nil, fmt.Errorf("failed to call contract: %w", err)
	}

	return contractABI.Unpack(method, result)
}

// SendTransaction sends a transaction to modify contract state
func (c *Client) SendTransaction(ctx context.Context, contractAddr common.Address, contractABI abi.ABI, from common.Address, privateKey *ecdsa.PrivateKey, method string, args ...interface{}) (*types.Transaction, error) {
	data, err := contractABI.Pack(method, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to pack method %s: %w", method, err)
	}

	// Get nonce
	nonce, err := c.ETHClient.PendingNonceAt(ctx, from)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}

	// Estimate gas using direct RPC call
	gasEstimateArgs := map[string]interface{}{
		"from": from,
		"to":   contractAddr,
		"data": data,
	}

	// Estimate gas with direct RPC call
	var gasLimit string
	err = c.ETHClient.Client().CallContext(ctx, &gasLimit, "eth_estimateGas", gasEstimateArgs)
	if err != nil {
		return nil, fmt.Errorf("failed to estimate gas: %w", err)
	}

	// Convert gas limit from hex string to uint64
	gasLimitInt, err := strconv.ParseUint(gasLimit[2:], 16, 64) // Remove 0x prefix
	if err != nil {
		return nil, fmt.Errorf("failed to parse gas limit: %w", err)
	}

	// Get gas price
	gasPrice, err := c.ETHClient.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get gas price: %w", err)
	}

	// Create transaction
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimitInt,
		To:       &contractAddr,
		Value:    big.NewInt(0),
		Data:     data,
	})

	// Get chain ID
	chainID, err := c.ETHClient.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

	// Sign transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Send transaction
	if err := c.ETHClient.SendTransaction(ctx, signedTx); err != nil {
		return nil, fmt.Errorf("failed to send transaction: %w", err)
	}

	return signedTx, nil
}
