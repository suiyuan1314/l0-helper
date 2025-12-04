package liquidity

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/l0-modular-arb-bot/config"
)

// Engine represents the Liquidity Engine for DEX interactions
type Engine struct {
	Config     *config.Config
	HTTPClient *http.Client
}

// NewEngine creates a new Liquidity Engine instance
type NewEngineParams struct {
	Config *config.Config
}

// NewEngine creates a new Liquidity Engine
func NewEngine(params NewEngineParams) (*Engine, error) {
	return &Engine{
		Config: params.Config,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}, nil
}

// SwapQuote represents a swap quote from the DEX aggregator
type SwapQuote struct {
	ToAddress     common.Address
	CallData      []byte
	ExpectedOut   *big.Int
	Gas           uint64
}

// GetSwapQuote fetches a swap quote from the DEX aggregator
type GetSwapQuoteParams struct {
	ChainName     string
	TokenIn       common.Address
	TokenOut      common.Address
	AmountIn      *big.Int
	Slippage      float64
}

// GetSwapQuote fetches a swap quote from the DEX aggregator
func (e *Engine) GetSwapQuote(ctx context.Context, params GetSwapQuoteParams) (*SwapQuote, error) {
	// Get chain configuration
	chainConfig := e.Config.GetChainConfig(params.ChainName)
	if chainConfig == nil {
		return nil, fmt.Errorf("chain %s not found in config", params.ChainName)
	}

	// Prepare request to DEX aggregator
	url, err := e.buildSwapURL(chainConfig, params)
	if err != nil {
		return nil, err
	}

	// Make request
	resp, err := e.HTTPClient.Get(url.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get swap quote: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Parse response
	var quoteResp aggregatorResponse
	if err := json.Unmarshal(body, &quoteResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	// Convert to SwapQuote
	return e.parseSwapQuote(quoteResp)
}

// CheckApproval checks if the bot has sufficient approval for the aggregator spender
type CheckApprovalParams struct {
	ChainName     string
	TokenAddress  common.Address
	SpenderAddress common.Address
	OwnerAddress  common.Address
	Client        *ethclient.Client
}

// CheckApproval checks if the bot has sufficient approval for the aggregator spender
func (e *Engine) CheckApproval(ctx context.Context, params CheckApprovalParams) (bool, error) {
	// Get chain configuration
	chainConfig := e.Config.GetChainConfig(params.ChainName)
	if chainConfig == nil {
		return false, fmt.Errorf("chain %s not found in config", params.ChainName)
	}

	// Load ERC20 ABI
	erc20ABI, err := loadERC20ABI()
	if err != nil {
		return false, err
	}

	// Call allowance method
	data, err := erc20ABI.Pack("allowance", params.OwnerAddress, params.SpenderAddress)
	if err != nil {
		return false, fmt.Errorf("failed to pack allowance method: %w", err)
	}

	// Use direct RPC call to eth_call
	callArgs := map[string]interface{}{
		"to":   params.TokenAddress,
		"data": data,
	}

	var result []byte
	err = params.Client.Client().CallContext(ctx, &result, "eth_call", callArgs, "latest")
	if err != nil {
		return false, fmt.Errorf("failed to call allowance: %w", err)
	}

	// Unpack result
	var allowance *big.Int
	if err := erc20ABI.UnpackIntoInterface(&allowance, "allowance", result); err != nil {
		return false, fmt.Errorf("failed to unpack allowance: %w", err)
	}

	// Check if allowance is sufficient (infinite approval is typically a very large number)
	return allowance.Cmp(big.NewInt(0)) > 0 && allowance.Cmp(big.NewInt(1000000000000000000)) > 0, nil
}

// Approve generates an approve transaction for the aggregator spender
type ApproveParams struct {
	ChainName     string
	TokenAddress  common.Address
	SpenderAddress common.Address
}

// Approve generates an approve transaction for the aggregator spender
func (e *Engine) Approve(ctx context.Context, params ApproveParams) ([]byte, common.Address, error) {
	// Load ERC20 ABI
	erc20ABI, err := loadERC20ABI()
	if err != nil {
		return nil, common.Address{}, err
	}

	// Generate approve call data with infinite approval
	data, err := erc20ABI.Pack("approve", params.SpenderAddress, big.NewInt(0).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1)))
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("failed to pack approve method: %w", err)
	}

	return data, params.TokenAddress, nil
}

// buildSwapURL builds the URL for the DEX aggregator API call
func (e *Engine) buildSwapURL(chainConfig *config.ChainConfig, params GetSwapQuoteParams) (*url.URL, error) {
	// Parse base URL
	baseURL, err := url.Parse(chainConfig.Aggregator.ApiUrl)
	if err != nil {
		return nil, fmt.Errorf("invalid aggregator API URL: %w", err)
	}

	// Set path for swap endpoint
	baseURL.Path = strings.TrimSuffix(baseURL.Path, "/") + "/swap"

	// Add query parameters
	q := baseURL.Query()
	q.Set("fromTokenAddress", params.TokenIn.Hex())
	q.Set("toTokenAddress", params.TokenOut.Hex())
	q.Set("amount", params.AmountIn.String())
	q.Set("slippage", fmt.Sprintf("%f", params.Slippage))
	q.Set("disableEstimate", "true")
	if chainConfig.Aggregator.ApiKey != "" {
		q.Set("apiKey", chainConfig.Aggregator.ApiKey)
	}
	baseURL.RawQuery = q.Encode()

	return baseURL, nil
}

// parseSwapQuote parses the aggregator response into a SwapQuote
func (e *Engine) parseSwapQuote(resp aggregatorResponse) (*SwapQuote, error) {
	// Extract relevant information from the response
	toAddress := common.HexToAddress(resp.Tx.To)
	callData, err := hex.DecodeString(strings.TrimPrefix(resp.Tx.Data, "0x"))
	if err != nil {
		return nil, fmt.Errorf("invalid call data: %w", err)
	}

	// Parse expected output amount
	expectedOut, ok := new(big.Int).SetString(resp.ToAmount, 10)
	if !ok {
		return nil, fmt.Errorf("invalid expected output amount: %s", resp.ToAmount)
	}

	// Parse gas estimate
	gas, err := strconv.ParseUint(resp.Tx.Gas, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid gas estimate: %w", err)
	}

	return &SwapQuote{
		ToAddress:   toAddress,
		CallData:    callData,
		ExpectedOut: expectedOut,
		Gas:         gas,
	}, nil
}

// aggregatorResponse represents the response from the DEX aggregator
type aggregatorResponse struct {
	ToAmount string `json:"toAmount"`
	Tx       struct {
		To    string `json:"to"`
		Data  string `json:"data"`
		Gas   string `json:"gas"`
	} `json:"tx"`
}

// loadERC20ABI loads the ERC20 ABI from file
func loadERC20ABI() (abi.ABI, error) {
	file, err := os.Open("abis/ERC20.json")
	if err != nil {
		return abi.ABI{}, fmt.Errorf("failed to open ERC20 ABI file: %w", err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return abi.ABI{}, fmt.Errorf("failed to read ERC20 ABI file: %w", err)
	}

	var erc20ABI abi.ABI
	if err := json.Unmarshal(content, &erc20ABI); err != nil {
		return abi.ABI{}, fmt.Errorf("failed to unmarshal ERC20 ABI: %w", err)
	}

	return erc20ABI, nil
}
