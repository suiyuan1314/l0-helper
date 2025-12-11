package contracts

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

// Client represents a contract client for interacting with LayerZero V2 contracts
type Client struct {
	ETHClient *ethclient.Client
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

	return &Client{
		ETHClient: client,
	}, nil
}
