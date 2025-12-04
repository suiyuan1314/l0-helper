package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Config represents the bot's configuration
type Config struct {
	Chains []ChainConfig `json:"chains"`
	Bot    BotConfig     `json:"bot"`
}

// ChainConfig represents configuration for a specific chain
type ChainConfig struct {
	Name       string         `json:"name"`
	Eid        uint32         `json:"eid"`
	ChainId    uint64         `json:"chainId"`
	RpcUrl     string         `json:"rpcUrl"`
	WsUrl      string         `json:"wsUrl"`
	Contracts  Contracts      `json:"contracts"`
	Aggregator Aggregator     `json:"aggregator"`
	Bundling   BundlingConfig `json:"bundling"`
}

// Contracts represents contract addresses for a chain
type Contracts struct {
	Endpoint         string `json:"endpoint"`
	EndpointView     string `json:"endpointView"`
	OFT              string `json:"oft"`
	AggregatorSpender string `json:"aggregatorSpender"`
}

// Aggregator represents DEX aggregator configuration
type Aggregator struct {
	ApiUrl  string `json:"apiUrl"`
	ApiKey  string `json:"apiKey"`
}

// BundlingConfig represents transaction bundling configuration
type BundlingConfig struct {
	Enabled  bool   `json:"enabled"`
	RelayUrl string `json:"relayUrl"`
	Mode     string `json:"mode"`
}

// BotConfig represents global bot configuration
type BotConfig struct {
	PrivateKey        string  `json:"privateKey"`
	MaxGapThreshold   uint64  `json:"maxGapThreshold"`
	FeeBufferMultiplier float64 `json:"feeBufferMultiplier"`
}

// LoadConfig loads the configuration from a JSON file
func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	// Validate config
	if err := config.Validate(); err != nil {
		return nil, err
	}

	// Process environment variables
	if config.Bot.PrivateKey == "ENV_VAR" {
		privateKey := os.Getenv("PRIVATE_KEY")
		if privateKey == "" {
			log.Fatal("PRIVATE_KEY environment variable not set")
		}
		config.Bot.PrivateKey = privateKey
	}

	return &config, nil
}

// Validate validates the configuration
func (c *Config) Validate() error {
	if len(c.Chains) == 0 {
		return &ValidationError{"no chains configured"}
	}

	for i, chain := range c.Chains {
		if chain.Name == "" {
			return &ValidationError{"chain name cannot be empty"}
		}
		if chain.RpcUrl == "" {
			return &ValidationError{"rpcUrl cannot be empty for chain " + chain.Name}
		}
		if chain.Contracts.Endpoint == "" {
			return &ValidationError{"endpoint contract cannot be empty for chain " + chain.Name}
		}
		if chain.Contracts.OFT == "" {
			return &ValidationError{"OFT contract cannot be empty for chain " + chain.Name}
		}
		if chain.Aggregator.ApiUrl == "" {
			return &ValidationError{"aggregator apiUrl cannot be empty for chain " + chain.Name}
		}
		c.Chains[i].Name = toLower(chain.Name)
	}

	if c.Bot.MaxGapThreshold == 0 {
		c.Bot.MaxGapThreshold = 3 // Default value
	}

	if c.Bot.FeeBufferMultiplier == 0 {
		c.Bot.FeeBufferMultiplier = 1.1 // Default value
	}

	return nil
}

// GetChainConfig returns the configuration for a specific chain by name
func (c *Config) GetChainConfig(chainName string) *ChainConfig {
	for _, chain := range c.Chains {
		if chain.Name == toLower(chainName) {
			return &chain
		}
	}
	return nil
}

// ValidationError represents a configuration validation error
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return "config validation error: " + e.Message
}

func toLower(s string) string {
	result := make([]byte, len(s))
	for i, c := range s {
		if c >= 'A' && c <= 'Z' {
			result[i] = byte(c + 32)
		} else {
			result[i] = byte(c)
		}
	}
	return string(result)
}
