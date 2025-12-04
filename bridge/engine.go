package bridge

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/l0-modular-arb-bot/config"
	"github.com/l0-modular-arb-bot/contracts"
)

// Engine represents the Bridge Engine for LayerZero V2 operations
type Engine struct {
	Config     *config.Config
}

// NewEngine creates a new Bridge Engine instance
type NewEngineParams struct {
	Config *config.Config
}

// NewEngine creates a new Bridge Engine
func NewEngine(params NewEngineParams) (*Engine, error) {
	return &Engine{
		Config: params.Config,
	}, nil
}

// SendResult represents the result of a send operation
type SendResult struct {
	Transaction *types.Transaction
	GUID        common.Hash
	Nonce       uint64
}

// SendParams represents parameters for sending a LayerZero message
type SendParams struct {
	ChainName       string
	From            common.Address
	ToAddress       common.Address
	Amount          *big.Int
	DstEID          uint32
	FeeBuffer       float64
}

// Send sends a LayerZero message from the source chain
func (e *Engine) Send(ctx context.Context, client *contracts.Client, params SendParams, privateKey *ecdsa.PrivateKey) (*SendResult, error) {
	// Get chain configuration
	chainConfig := e.Config.GetChainConfig(params.ChainName)
	if chainConfig == nil {
		return nil, fmt.Errorf("chain %s not found in config", params.ChainName)
	}

	// Convert receiver address to bytes32
	receiverBytes32 := common.BytesToHash(params.ToAddress.Bytes())

	// Call quoteSend to estimate fees
	quoteArgs := []interface{}{
		params.DstEID,
		receiverBytes32,
		params.Amount,
		false,
		[]byte{}, // options
	}

	quoteResult, err := client.CallContract(ctx, common.HexToAddress(chainConfig.Contracts.OFT), client.OFTABI, "quoteSend", quoteArgs...)
	if err != nil {
		return nil, fmt.Errorf("failed to call quoteSend: %w", err)
	}

	// Parse quote result
	if len(quoteResult) != 1 {
		return nil, fmt.Errorf("invalid quoteSend result length: %d", len(quoteResult))
	}

	nativeFee, ok := quoteResult[0].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("failed to parse nativeFee from quoteSend")
	}

	// Apply fee buffer
	feeWithBuffer := new(big.Int).Mul(nativeFee, big.NewInt(int64(params.FeeBuffer*1000000000)))
	feeWithBuffer = feeWithBuffer.Div(feeWithBuffer, big.NewInt(1000000000))

	// Build SendParam struct
	sendParam := struct {
		dstEid     uint32
		receiver   common.Address
		amount     *big.Int
		minAmount  *big.Int
		lzReceiveOption []byte
		extraOptions     []byte
		composer        common.Address
	}{}

	// Set SendParam fields
	sendParam.dstEid = params.DstEID
	sendParam.receiver = params.ToAddress
	sendParam.amount = params.Amount
	sendParam.minAmount = params.Amount // Assuming full amount is expected
	sendParam.lzReceiveOption = e.buildLzReceiveOption()
	sendParam.extraOptions = []byte{}
	sendParam.composer = common.Address{}

	// Pack SendParam into bytes directly in SendTransaction call

	// Send transaction
	tx, err := client.SendTransaction(ctx, common.HexToAddress(chainConfig.Contracts.OFT), client.OFTABI, params.From, privateKey, "send", sendParam)
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %w", err)
	}

	// Return result
	return &SendResult{
		Transaction: tx,
		GUID:        common.Hash{}, // TODO: Extract GUID from event
		Nonce:       0,              // TODO: Extract nonce from event
	}, nil
}

// ReceiveResult represents the result of a receive operation
type ReceiveResult struct {
	Transaction *types.Transaction
	Amount      *big.Int
}

// ReceiveParams represents parameters for receiving a LayerZero message
type ReceiveParams struct {
	ChainName       string
	From            common.Address
	Origin          Origin
	GUID            common.Hash
	Message         []byte
	ExtraData       []byte
}

// Origin represents the origin of a LayerZero message
type Origin struct {
	SrcEid  uint32
	Sender  common.Hash
	Nonce   uint64
}

// Receive processes a LayerZero message on the destination chain
func (e *Engine) Receive(ctx context.Context, client *contracts.Client, params ReceiveParams, privateKey *ecdsa.PrivateKey) (*ReceiveResult, error) {
	// Get chain configuration
	chainConfig := e.Config.GetChainConfig(params.ChainName)
	if chainConfig == nil {
		return nil, fmt.Errorf("chain %s not found in config", params.ChainName)
	}

	// No need to build data explicitly - SendTransaction will handle packing

	// Send transaction
	tx, err := client.SendTransaction(ctx, common.HexToAddress(chainConfig.Contracts.Endpoint), client.EndpointV2ABI, params.From, privateKey, "lzReceive", 
		params.Origin,
		params.From, // receiver
		params.GUID,
		params.Message,
		params.ExtraData,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to send lzReceive transaction: %w", err)
	}

	// Decode amount from message for return
	amount, err := e.decodeOFTAmount(params.Message)
	if err != nil {
		return nil, fmt.Errorf("failed to decode OFT amount: %w", err)
	}

	return &ReceiveResult{
		Transaction: tx,
		Amount:      amount,
	}, nil
}

// decodeOFTAmount decodes the amount from an OFT message
func (e *Engine) decodeOFTAmount(message []byte) (*big.Int, error) {
	// OFT payload layout: Receiver Address (32 bytes) + AmountSD (8 bytes)
	if len(message) < 40 {
		return nil, fmt.Errorf("invalid OFT message length: %d", len(message))
	}

	// Extract AmountSD (uint64) from bytes 32-40
	amountSDBytes := message[32:40]
	amountSD := new(big.Int).SetBytes(amountSDBytes)

	// Convert AmountSD to AmountLD using decimal conversion formula
	// Formula: AmountLD = AmountSD * 10^(LocalDecimals - SharedDecimals)
	// SharedDecimals = 6, LocalDecimals = 18
	conversionFactor := new(big.Int).Exp(big.NewInt(10), big.NewInt(12), nil) // 10^(18-6) = 10^12
	amountLD := new(big.Int).Mul(amountSD, conversionFactor)

	return amountLD, nil
}

// buildLzReceiveOption builds the lzReceiveOption bytes for SendParam
func (e *Engine) buildLzReceiveOption() []byte {
	// Options Type 3: TYPE_3 (0x0003) + WORKER_ID (0x01) + OPTION_LENGTH + OPTION_TYPE + PARAMS
	// Option Type 1: gasLimit (uint128) + msgValue (uint128)
	var options []byte

	// Type 3 header
	options = append(options, 0x00, 0x03) // TYPE_3
	options = append(options, 0x01)       // WORKER_ID

	// Option length placeholder (will be filled later)
	options = append(options, 0x00, 0x00) // OPTION_LENGTH placeholder

	// Option Type 1
	options = append(options, 0x01) // OPTION_TYPE

	// gasLimit (uint128) - set to a reasonable value for lzReceive
	gasLimit := big.NewInt(500000) // 500k gas
	gasLimitBytes := make([]byte, 16)
	gasLimit.FillBytes(gasLimitBytes)
	options = append(options, gasLimitBytes...)

	// msgValue (uint128) - set to 0 since we're not sending value
	msgValueBytes := make([]byte, 16)
	options = append(options, msgValueBytes...)

	// Update option length
	optionLength := uint16(len(options) - 4) // subtract header length
	options[2] = byte(optionLength >> 8)
	options[3] = byte(optionLength)

	return options
}

// MonitorPacketVerified monitors PacketVerified events on a chain
type MonitorPacketVerifiedParams struct {
	ChainName   string
	Client      *ethclient.Client
	Handler     func(event PacketVerifiedEvent)
}

// PacketVerifiedEvent represents a PacketVerified event
type PacketVerifiedEvent struct {
	Origin   Origin
	Receiver common.Address
	GUID     common.Hash
}

// MonitorPacketVerified monitors PacketVerified events on a chain
func (e *Engine) MonitorPacketVerified(ctx context.Context, params MonitorPacketVerifiedParams) error {
	// Get chain configuration
	chainConfig := e.Config.GetChainConfig(params.ChainName)
	if chainConfig == nil {
		return fmt.Errorf("chain %s not found in config", params.ChainName)
	}

	// Load EndpointV2 ABI
	endpointV2ABI, err := loadEndpointV2ABI()
	if err != nil {
		return err
	}

	// Get event ID for PacketVerified
	eventID := endpointV2ABI.Events["PacketVerified"].ID

	// Create event filter query using ethereum.FilterQuery
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(chainConfig.Contracts.Endpoint)},
		Topics:    [][]common.Hash{{eventID}},
	}

	// Create event channel
	events := make(chan types.Log)
	sub, err := params.Client.SubscribeFilterLogs(ctx, query, events)
	if err != nil {
		return fmt.Errorf("failed to subscribe to PacketVerified events: %w", err)
	}
	defer sub.Unsubscribe()

	// Process events
	for {
		select {
		case err := <-sub.Err():
			return fmt.Errorf("subscription error: %w", err)
		case vLog := <-events:
			// Parse the event
			event, err := e.parsePacketVerifiedEvent(endpointV2ABI, vLog)
			if err != nil {
				fmt.Printf("failed to parse PacketVerified event: %v\n", err)
				continue
			}

			// Call the handler
			params.Handler(event)
		case <-ctx.Done():
			return nil
		}
	}
}

// parsePacketVerifiedEvent parses a PacketVerified event from a log
func (e *Engine) parsePacketVerifiedEvent(endpointV2ABI abi.ABI, vLog types.Log) (PacketVerifiedEvent, error) {
	var event PacketVerifiedEvent

	// Unpack the event data
	err := endpointV2ABI.UnpackIntoInterface(&event, "PacketVerified", vLog.Data)
	if err != nil {
		return event, fmt.Errorf("failed to unpack PacketVerified event: %w", err)
	}

	// Extract topics (Origin struct)
	if len(vLog.Topics) < 4 {
		return event, fmt.Errorf("invalid PacketVerified event topic count: %d", len(vLog.Topics))
	}

	// Parse Origin struct from topics
	event.Origin.SrcEid = uint32(new(big.Int).SetBytes(vLog.Topics[1].Bytes()).Uint64())
	event.Origin.Sender = vLog.Topics[2]
	event.Origin.Nonce = new(big.Int).SetBytes(vLog.Topics[3].Bytes()).Uint64()
	event.GUID = vLog.Topics[0]

	return event, nil
}

// loadEndpointV2ABI loads the EndpointV2 ABI from file
func loadEndpointV2ABI() (abi.ABI, error) {
	file, err := os.Open("abis/EndpointV2.json")
	if err != nil {
		return abi.ABI{}, fmt.Errorf("failed to open EndpointV2 ABI file: %w", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return abi.ABI{}, fmt.Errorf("failed to read EndpointV2 ABI file: %w", err)
	}

	var endpointV2ABI abi.ABI
	if err := json.Unmarshal(content, &endpointV2ABI); err != nil {
		return abi.ABI{}, fmt.Errorf("failed to unmarshal EndpointV2 ABI: %w", err)
	}

	return endpointV2ABI, nil
}
