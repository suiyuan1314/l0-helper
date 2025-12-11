package liquidity

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/l0-modular-arb-bot/config"
	"github.com/l0-modular-arb-bot/contracts/erc20"
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
			Timeout: 20 * time.Second, // Increased timeout for 2-step process
		},
	}, nil
}

// SwapQuote represents a swap quote from the DEX aggregator
type SwapQuote struct {
	ToAddress   common.Address
	CallData    []byte
	ExpectedOut *big.Int
	Gas         uint64
}

// GetSwapQuoteParams represents parameters for fetching a swap quote
type GetSwapQuoteParams struct {
	ChainName string
	TokenIn   common.Address
	TokenOut  common.Address
	AmountIn  *big.Int
	Slippage  float64
	Sender    common.Address // Added sender for KyberSwap
	Recipient common.Address // Added recipient for KyberSwap
}

// KyberSwap V1 API Structs

type kyberGetRouteResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		RouteSummary  json.RawMessage `json:"routeSummary"` // Keep raw for passing to build
		RouterAddress string          `json:"routerAddress"`
	} `json:"data"`
}

type kyberBuildRouteRequest struct {
	RouteSummary json.RawMessage `json:"routeSummary"`
	Sender       string          `json:"sender"`
	Recipient    string          `json:"recipient"`
	Slippage     float64         `json:"slippageTolerance"`
}

type kyberBuildRouteResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		AmountOut     string `json:"amountOut"`
		Gas           string `json:"gas"`
		Data          string `json:"data"`
		RouterAddress string `json:"routerAddress"`
	} `json:"data"`
}

// GetSwapQuote fetches a swap quote from KyberSwap V1 API
func (e *Engine) GetSwapQuote(ctx context.Context, params GetSwapQuoteParams) (*SwapQuote, error) {
	// Get chain configuration
	chainConfig := e.Config.GetChainConfig(params.ChainName)
	if chainConfig == nil {
		return nil, fmt.Errorf("chain %s not found in config", params.ChainName)
	}

	// 1. Get Route (GET /routes)
	routeURL, err := e.buildRouteURL(chainConfig, params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "GET", routeURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create route request: %w", err)
	}
	req.Header.Set("x-client-id", "L0ModularArbBot")

	resp, err := e.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get swap route: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("get routes failed with status %d: %s", resp.StatusCode, string(body))
	}

	var routeResp kyberGetRouteResponse
	if err := json.NewDecoder(resp.Body).Decode(&routeResp); err != nil {
		return nil, fmt.Errorf("failed to decode route response: %w", err)
	}

	if routeResp.Code != 0 && routeResp.Code != 200 { // KyberSwap might use 0 or 200 for success
		return nil, fmt.Errorf("kyberswap route error: %s", routeResp.Message)
	}

	// 2. Build Route (POST /route/build)
	buildURL, err := e.buildBuildURL(chainConfig)
	if err != nil {
		return nil, err
	}

	// Prepare build request body
	buildReqBody := kyberBuildRouteRequest{
		RouteSummary: routeResp.Data.RouteSummary,
		Sender:       params.Sender.Hex(),
		Recipient:    params.Recipient.Hex(),
		Slippage:     params.Slippage * 100, // KyberSwap uses bps (10 = 0.1%), input is usually percentage (0.5 for 0.5%)?
		// Wait, doc says: "The value is in ranges [0, 2000], with 10 meaning 0.1%, and 0.1 meaning 0.001%."
		// If params.Slippage is e.g. 0.5 (meaning 0.5%), then 0.5% = 50 bps.
		// If params.Slippage is 0.005 (meaning 0.5%), then 0.5% = 50 bps.
		// Let's assume params.Slippage is stored as percentage float (e.g. 0.5 for 0.5%).
		// So 0.5 * 100 = 50 bps.
	}
	// Correcting slippage calculation based on doc: "The unit is bps (1/100 of 1%). ... 10 meaning 0.1%"
	// So 1 bps = 0.01%. 10 bps = 0.1%.
	// If params.Slippage is 0.5 (percent), we need 50.
	// So params.Slippage * 100 seems correct if params.Slippage is the number (0.5) and not fraction (0.005).
	// Assuming params.Slippage is like 0.5 for 0.5%.

	buildBody, err := json.Marshal(buildReqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal build request: %w", err)
	}

	reqBuild, err := http.NewRequestWithContext(ctx, "POST", buildURL.String(), bytes.NewBuffer(buildBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create build request: %w", err)
	}
	reqBuild.Header.Set("Content-Type", "application/json")
	reqBuild.Header.Set("x-client-id", "L0ModularArbBot")

	respBuild, err := e.HTTPClient.Do(reqBuild)
	if err != nil {
		return nil, fmt.Errorf("failed to build swap route: %w", err)
	}
	defer respBuild.Body.Close()

	if respBuild.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(respBuild.Body)
		return nil, fmt.Errorf("build route failed with status %d: %s", respBuild.StatusCode, string(body))
	}

	var buildResp kyberBuildRouteResponse
	if err := json.NewDecoder(respBuild.Body).Decode(&buildResp); err != nil {
		return nil, fmt.Errorf("failed to decode build response: %w", err)
	}

	if buildResp.Code != 0 && buildResp.Code != 200 {
		return nil, fmt.Errorf("kyberswap build error: %s", buildResp.Message)
	}

	return e.parseKyberSwapQuote(buildResp)
}

// CheckApproval checks if the bot has sufficient approval for the aggregator spender
type CheckApprovalParams struct {
	ChainName      string
	TokenAddress   common.Address
	SpenderAddress common.Address
	OwnerAddress   common.Address
	RequiredAmount *big.Int
	Client         *ethclient.Client
}

// CheckApproval checks if the bot has sufficient approval for the aggregator spender
func (e *Engine) CheckApproval(ctx context.Context, params CheckApprovalParams) (bool, error) {
	erc20Instance, err := erc20.NewERC20(params.TokenAddress, params.Client)
	if err != nil {
		return false, fmt.Errorf("failed to create ERC20 binding: %w", err)
	}

	// Call allowance method
	allowance, err := erc20Instance.Allowance(&bind.CallOpts{Context: ctx}, params.OwnerAddress, params.SpenderAddress)
	if err != nil {
		return false, fmt.Errorf("failed to call allowance: %w", err)
	}

	// Check if allowance is sufficient
	return allowance.Cmp(params.RequiredAmount) >= 0, nil
}

// Approve generates an approve transaction for the aggregator spender
type ApproveParams struct {
	ChainName      string
	TokenAddress   common.Address
	SpenderAddress common.Address
}

// Approve generates an approve transaction for the aggregator spender
func (e *Engine) Approve(ctx context.Context, params ApproveParams) ([]byte, common.Address, error) {
	// Get ABI
	erc20ABI, err := erc20.ERC20MetaData.GetAbi()
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("failed to get ERC20 ABI: %w", err)
	}

	// Generate approve call data with infinite approval
	// MaxUint256
	maxUint256 := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))
	data, err := erc20ABI.Pack("approve", params.SpenderAddress, maxUint256)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("failed to pack approve method: %w", err)
	}

	return data, params.TokenAddress, nil
}

// buildRouteURL builds the URL for the KyberSwap GET /routes API call
func (e *Engine) buildRouteURL(chainConfig *config.ChainConfig, params GetSwapQuoteParams) (*url.URL, error) {
	// Parse base URL (should be like https://aggregator-api.kyberswap.com/chain/api)
	// We want https://aggregator-api.kyberswap.com/chain/api/v1/routes
	baseURL, err := url.Parse(chainConfig.Aggregator.ApiUrl)
	if err != nil {
		return nil, fmt.Errorf("invalid aggregator API URL: %w", err)
	}

	// Append /v1/routes
	// Ensure no double slashes
	baseURL.Path = strings.TrimSuffix(baseURL.Path, "/") + "/v1/routes"

	// Add query parameters
	q := baseURL.Query()
	q.Set("tokenIn", params.TokenIn.Hex())
	q.Set("tokenOut", params.TokenOut.Hex())
	q.Set("amountIn", params.AmountIn.String())
	// Note: slippage is not needed for GET route in V1 unless filtering,
	// but the doc says "find best route". It doesn't list slippage as required for GET.
	// It is required for POST.

	// q.Set("gasInclude", "true") // Defaults to true

	baseURL.RawQuery = q.Encode()

	return baseURL, nil
}

// buildBuildURL builds the URL for the KyberSwap POST /route/build API call
func (e *Engine) buildBuildURL(chainConfig *config.ChainConfig) (*url.URL, error) {
	// Parse base URL
	baseURL, err := url.Parse(chainConfig.Aggregator.ApiUrl)
	if err != nil {
		return nil, fmt.Errorf("invalid aggregator API URL: %w", err)
	}

	// Append /v1/route/build
	baseURL.Path = strings.TrimSuffix(baseURL.Path, "/") + "/v1/route/build"
	return baseURL, nil
}

// parseKyberSwapQuote parses the KyberSwap build response into a SwapQuote
func (e *Engine) parseKyberSwapQuote(resp kyberBuildRouteResponse) (*SwapQuote, error) {
	// Extract relevant information from the response
	toAddress := common.HexToAddress(resp.Data.RouterAddress)

	callData, err := hex.DecodeString(strings.TrimPrefix(resp.Data.Data, "0x"))
	if err != nil {
		return nil, fmt.Errorf("invalid call data: %w", err)
	}

	// Parse expected output amount
	expectedOut, ok := new(big.Int).SetString(resp.Data.AmountOut, 10)
	if !ok {
		return nil, fmt.Errorf("invalid expected output amount: %s", resp.Data.AmountOut)
	}

	// Parse gas estimate
	gas, err := strconv.ParseUint(resp.Data.Gas, 10, 64)
	if err != nil {
		// Log warning?
		gas = 500000 // default or 0
	}

	return &SwapQuote{
		ToAddress:   toAddress,
		CallData:    callData,
		ExpectedOut: expectedOut,
		Gas:         gas,
	}, nil
}
