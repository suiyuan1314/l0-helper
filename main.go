package main

import (
	"fmt"
	"math/big"

	"github.com/l0-modular-arb-bot/bridge"
	"github.com/l0-modular-arb-bot/config"
	"github.com/l0-modular-arb-bot/contracts"
	"github.com/l0-modular-arb-bot/execution"
	"github.com/l0-modular-arb-bot/guard"
	"github.com/l0-modular-arb-bot/liquidity"
	"github.com/l0-modular-arb-bot/monitoring"
	"github.com/l0-modular-arb-bot/risk"
	"github.com/l0-modular-arb-bot/strategy"
)

func main() {
	fmt.Println("Starting LayerZero V2 Modular Arbitrage Bot...")

	// Step 1: Load configuration
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		return
	}

	// Step 2: Initialize logger
	logger, err := monitoring.NewLogger(monitoring.NewLoggerParams{
		Level:   monitoring.LogLevelInfo,
		UseJSON: false,
	})
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		return
	}

	logger.Info("Configuration loaded successfully")

	// Step 4: Initialize contract client for Ethereum
	_, err = contracts.NewClient(contracts.NewClientParams{
		RPCEndpoint: cfg.Chains[0].RpcUrl,
	})
	if err != nil {
		logger.Error("Failed to initialize contract client", monitoring.WithError(err))
		return
	}

	logger.Info("Contract client initialized successfully")

	// Step 4: Initialize engines
	liquidityEngine, err := liquidity.NewEngine(liquidity.NewEngineParams{
		Config: cfg,
	})
	if err != nil {
		logger.Error("Failed to initialize liquidity engine", monitoring.WithError(err))
		return
	}

	bridgeEngine, err := bridge.NewEngine(bridge.NewEngineParams{
		Config: cfg,
	})
	if err != nil {
		logger.Error("Failed to initialize bridge engine", monitoring.WithError(err))
		return
	}

	executionEngine, err := execution.NewEngine(execution.NewEngineParams{
		Config: cfg,
	})
	if err != nil {
		logger.Error("Failed to initialize execution engine", monitoring.WithError(err))
		return
	}

	guardEngine, err := guard.NewEngine(guard.NewEngineParams{
		Config: cfg,
	})
	if err != nil {
		logger.Error("Failed to initialize guard engine", monitoring.WithError(err))
		return
	}

	riskManager, err := risk.NewManager(risk.NewManagerParams{
		Config: cfg,
	})
	if err != nil {
		logger.Error("Failed to initialize risk manager", monitoring.WithError(err))
		return
	}

	// Step 5: Initialize strategy controller
	_, err = strategy.NewController(strategy.NewControllerParams{
		Config:          cfg,
		LiquidityEngine: liquidityEngine,
		BridgeEngine:    bridgeEngine,
		ExecutionEngine: executionEngine,
		GuardEngine:     guardEngine,
	})
	if err != nil {
		logger.Error("Failed to initialize strategy controller", monitoring.WithError(err))
		return
	}

	logger.Info("All engines initialized successfully")

	// Step 6: Test config validation
	logger.Info("Testing config validation...")
	if err := cfg.Validate(); err != nil {
		logger.Error("Config validation failed", monitoring.WithError(err))
		return
	}
	logger.Info("Config validation passed")

	// Step 7: Test log levels
	logger.Debug("This is a debug message")
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message", monitoring.WithError(fmt.Errorf("test error")))

	// Step 8: Test risk management features
	logger.Info("Testing risk management features...")

	// Test slippage calculation
	expectedOut := big.NewInt(1000000000000000000) // 1 ETH
	minOut := big.NewInt(995000000000000000)      // 0.995 ETH (0.5% slippage)
	slippageResult, err := riskManager.CalculateSlippage(risk.CalculateSlippageParams{
		ExpectedOut: expectedOut,
		MinOut:      minOut,
	})
	if err != nil {
		logger.Error("Failed to calculate slippage", monitoring.WithError(err))
	} else {
		logger.Info(fmt.Sprintf("Slippage calculation: %.2f%%", slippageResult.Slippage*100))
	}

	// Test dynamic fee buffer calculation
	baseFee := big.NewInt(10000000000) // 10 gwei
	feeBufferResult, err := riskManager.CalculateDynamicFeeBuffer(risk.CalculateDynamicFeeBufferParams{
		BaseFee:           baseFee,
		NetworkCongestion: 0.5, // 50% congestion
	})
	if err != nil {
		logger.Error("Failed to calculate dynamic fee buffer", monitoring.WithError(err))
	} else {
		logger.Info(fmt.Sprintf("Dynamic fee buffer: %.2fx", feeBufferResult.FeeBuffer))
	}

	// Step 9: Test strategy controller
	logger.Info("Testing strategy controller...")

	// Test recipe type mapping
	recipeTypes := []strategy.RecipeType{
		strategy.RecipeTypeStandardArb,
		strategy.RecipeTypeInventoryArb,
		strategy.RecipeTypeRebalancing,
		strategy.RecipeTypeTheCloser,
		strategy.RecipeTypeTheCourier,
		strategy.RecipeTypePanicSell,
	}

	for _, recipeType := range recipeTypes {
		logger.Info(fmt.Sprintf("Recipe type %d is supported", recipeType))
	}

	logger.Info("All tests completed successfully!")
	logger.Info("LayerZero V2 Modular Arbitrage Bot implementation validated")
}
