package bridge

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/l0-modular-arb-bot/config"
	"github.com/l0-modular-arb-bot/contracts"
	"github.com/l0-modular-arb-bot/contracts/endpointv2"
	"github.com/l0-modular-arb-bot/contracts/ioft"
)

// Engine represents the Bridge Engine for LayerZero V2 operations
type Engine struct {
	Config *config.Config
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
	ChainName string
	From      common.Address
	ToAddress common.Address
	Amount    *big.Int
	DstEID    uint32
	FeeBuffer float64
}

// Send sends a LayerZero message from the source chain
func (e *Engine) Send(ctx context.Context, client *contracts.Client, params SendParams, privateKey *ecdsa.PrivateKey) (*SendResult, error) {
	// Get chain configuration
	chainConfig := e.Config.GetChainConfig(params.ChainName)
	if chainConfig == nil {
		return nil, fmt.Errorf("chain %s not found in config", params.ChainName)
	}

	oftAddr := common.HexToAddress(chainConfig.Contracts.OFT)
	ioftInstance, err := ioft.NewIOFT(oftAddr, client.ETHClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create IOFT binding: %w", err)
	}

	receiverBytes32 := common.BytesToHash(params.ToAddress.Bytes())

	// Build SendParam
	sendParam := ioft.SendParam{
		DstEid:       params.DstEID,
		To:           receiverBytes32,
		AmountLD:     params.Amount,
		MinAmountLD:  params.Amount,            // Assuming 100% fixed for now, practically should imply slippage
		ExtraOptions: e.buildLzReceiveOption(), // Using legacy function for now, simpler
		ComposeMsg:   []byte{},
		OftCmd:       []byte{},
	}

	// Quote Send
	// quoteSend(SendParam calldata _sendParam, bool _payInLzToken) external view returns (MessagingFee memory msgFee);
	msgFee, err := ioftInstance.QuoteSend(&bind.CallOpts{Context: ctx}, sendParam, false)
	if err != nil {
		return nil, fmt.Errorf("failed to quote send: %w", err)
	}

	// Apply fee buffer to NativeFee
	nativeFee := msgFee.NativeFee
	if params.FeeBuffer > 0 {
		bufferMult := big.NewInt(int64((1.0 + params.FeeBuffer) * 100))
		nativeFee = new(big.Int).Mul(nativeFee, bufferMult)
		nativeFee = nativeFee.Div(nativeFee, big.NewInt(100))
	}
	msgFee.NativeFee = nativeFee

	// Prepare transaction options
	chainID, err := client.ETHClient.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %w", err)
	}
	auth.Value = nativeFee // CRITICAL FIX: Send native fee with transaction
	auth.Context = ctx

	// Send Transaction
	// send(SendParam calldata _sendParam, MessagingFee calldata _fee, address _refundAddress) external payable returns (MessagingReceipt memory msgReceipt, OFTReceipt memory oftReceipt);
	tx, err := ioftInstance.Send(auth, sendParam, msgFee, params.From)
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %w", err)
	}

	return &SendResult{
		Transaction: tx,
		GUID:        common.Hash{}, // TODO: Parse logs to get GUID if strictly needed immediately, or wait for event
		Nonce:       0,
	}, nil
}

// ReceiveResult represents the result of a receive operation
type ReceiveResult struct {
	Transaction *types.Transaction
	Amount      *big.Int
}

// ReceiveParams represents parameters for receiving a LayerZero message
type ReceiveParams struct {
	ChainName string
	From      common.Address
	Origin    endpointv2.Origin
	GUID      common.Hash
	Message   []byte
	ExtraData []byte
}

// Receive processes a LayerZero message on the destination chain
func (e *Engine) Receive(ctx context.Context, client *contracts.Client, params ReceiveParams, privateKey *ecdsa.PrivateKey) (*ReceiveResult, error) {
	// Get chain configuration
	chainConfig := e.Config.GetChainConfig(params.ChainName)
	if chainConfig == nil {
		return nil, fmt.Errorf("chain %s not found in config", params.ChainName)
	}

	endpointAddr := common.HexToAddress(chainConfig.Contracts.Endpoint)
	endpoint, err := endpointv2.NewEndpointV2(endpointAddr, client.ETHClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create Endpoint binding: %w", err)
	}

	// Prepare auth
	chainID, err := client.ETHClient.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %w", err)
	}
	auth.Context = ctx

	// lzReceive(Origin calldata _origin, address _receiver, bytes32 _guid, bytes calldata _message, bytes calldata _extraData) external payable;
	tx, err := endpoint.LzReceive(auth, params.Origin, params.From, params.GUID, params.Message, params.ExtraData)
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
	ChainName string
	Client    *ethclient.Client
	Handler   func(event PacketVerifiedEvent)
}

// PacketVerifiedEvent represents a PacketVerified event
type PacketVerifiedEvent struct {
	Origin      endpointv2.Origin
	Receiver    common.Address
	PayloadHash [32]byte
}

// MonitorPacketVerified monitors PacketVerified events on a chain
func (e *Engine) MonitorPacketVerified(ctx context.Context, params MonitorPacketVerifiedParams) error {
	// Get chain configuration
	chainConfig := e.Config.GetChainConfig(params.ChainName)
	if chainConfig == nil {
		return fmt.Errorf("chain %s not found in config", params.ChainName)
	}

	endpointAddr := common.HexToAddress(chainConfig.Contracts.Endpoint)
	filterer, err := endpointv2.NewEndpointV2Filterer(endpointAddr, params.Client)
	if err != nil {
		return fmt.Errorf("failed to create endpoint filterer: %w", err)
	}

	sink := make(chan *endpointv2.EndpointV2PacketVerified)
	sub, err := filterer.WatchPacketVerified(&bind.WatchOpts{Context: ctx}, sink)
	if err != nil {
		return fmt.Errorf("failed to watch PacketVerified events: %w", err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			return fmt.Errorf("subscription error: %w", err)
		case event := <-sink:
			params.Handler(PacketVerifiedEvent{
				Origin:      event.Origin,
				Receiver:    event.Receiver,
				PayloadHash: event.PayloadHash,
			})
		case <-ctx.Done():
			return nil
		}
	}
}
