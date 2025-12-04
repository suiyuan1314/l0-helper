# LayerZero V2 Modular Arbitrage Bot

A modular, high-frequency arbitrage (HFA) bot framework built for LayerZero V2, following a "Lego Block" architecture. This bot can dynamically compose standardized modules to form various arbitrage strategies across EVM-compatible chains.

## 📋 Prerequisites

Before you start, make sure you have:

1. **Go 1.25.4 or newer** - [Download Go](https://go.dev/dl/)
2. **A code editor** - VS Code (recommended), GoLand, or any text editor
3. **Git** - To clone the repository
4. **Ethereum Wallet** - With testnet ETH/stablecoins for testing
5. **RPC Endpoints** - For the chains you want to arbitrage on (Alchemy, Infura, or public RPCs)

## 📂 Project Structure

The bot uses a **modular "Lego Block" architecture** where each component (Engine) works independently and can be combined to create different strategies:

```
├── abis/               # Contract ABIs for LayerZero V2
├── bridge/             # LayerZero bridging operations (send/receive)
├── config/             # Configuration management
├── contracts/          # Smart contract interactions
├── execution/          # Transaction management with bundling support
├── guard/              # Risk validation and nonce gap handling
├── liquidity/          # DEX aggregator interactions for optimal swap routes
├── monitoring/         # Logging and monitoring
├── risk/               # Advanced risk management
├── strategy/           # Orchestrates all six arbitrage recipes
├── config.json         # Bot configuration file
├── go.mod              # Go dependencies
└── main.go             # Entry point for testing
```

## 🚀 Getting Started

### 1. Clone the Repository

Open your terminal and clone this repository:

```bash
git clone https://github.com/your-username/l0-modular-arb-bot.git
cd l0-modular-arb-bot
```

### 2. Install Dependencies

Go modules manage dependencies automatically. Run this command to download all required packages:

```bash
go mod tidy
```

### 3. Configure the Bot

Edit `config.json` to set up your chains and preferences. Here's a breakdown of the configuration:

```json
{
  "chains": [
    {
      "name": "Ethereum",
      "eid": 30101,              // LayerZero chain ID
      "chainId": 1,             // Ethereum mainnet chain ID (use 5 for Goerli testnet)
      "rpcUrl": "https://eth-mainnet.g.alchemy.com/v2/YOUR_KEY",
      "wsUrl": "wss://eth-mainnet.g.alchemy.com/v2/YOUR_KEY",
      "contracts": {
        "endpoint": "0x1a44076050125825900e736c501f859c50fE728c",  // LayerZero EndpointV2
        "endpointView": "",    // EndpointV2View contract
        "oft": "0x...",        // OFT token address
        "aggregatorSpender": "0x1111111254fb6c44bac0bed2854e76f90643097d"  // DEX aggregator spender
      },
      "aggregator": {
        "apiUrl": "https://api.1inch.dev/swap/v6.0/1/swap",  // DEX aggregator API
        "apiKey": "YOUR_1INCH_KEY"
      },
      "bundling": {
        "enabled": false,       // Enable Flashbots/Jito bundling
        "relayUrl": "https://relay.flashbots.net"
      }
    }
  ],
  "bot": {
    "privateKey": "ENV_VAR",    // Use environment variable for private key
    "maxGapThreshold": 3,       // Maximum nonce gap to resolve
    "feeBufferMultiplier": 1.1  // Safety buffer for LayerZero fees
  }
}
```

**Quick Tips for Beginners**:
- Start with testnets (e.g., Goerli for Ethereum, BSC Testnet)
- Get free testnet tokens from faucets like [Goerli Faucet](https://goerlifaucet.com/)
- Use public RPCs if you don't have Alchemy/Infura keys

### 4. Set Up Environment Variables

Never hardcode your private key! Use an environment variable instead:

**Mac/Linux**:
```bash
export PRIVATE_KEY="your-private-key-here"
```

**Windows (Command Prompt)**:
```cmd
set PRIVATE_KEY=your-private-key-here
```

**Windows (PowerShell)**:
```powershell
$env:PRIVATE_KEY="your-private-key-here"
```

## 🔧 Testing the Bot

Run the test script to verify everything works correctly:

```bash
go run main.go
```

You should see output like:
```
Starting LayerZero V2 Modular Arbitrage Bot...
2025/12/03 14:30:00 INFO Configuration loaded successfully
2025/12/03 14:30:00 INFO Contract client initialized successfully
2025/12/03 14:30:00 INFO All engines initialized successfully
2025/12/03 14:30:00 INFO Config validation passed
...
All tests completed successfully!
LayerZero V2 Modular Arbitrage Bot implementation validated
```

## 🍲 Arbitrage Recipes

The bot supports 6 arbitrage strategies ("recipes"):

1. **Standard Arb**: Buy (Chain A) → Bridge → Sell (Chain B)
   *Use case*: Capital is in ETH/stables; opportunity detected across chains

2. **Inventory Arb**: Bridge (Chain A) → Sell (Chain B)
   *Use case*: Bot already holds OFT on the source chain

3. **Rebalancing**: Bridge (Chain A) → Receive (Chain B)
   *Use case*: Move inventory to a chain with better future opportunities

4. **The Closer**: Receive (Chain B) → Sell (Chain B)
   *Use case*: Finalize an in-flight arbitrage or "catch" a stuck packet

5. **The Courier**: Pure Receive (Chain B)
   *Use case*: Claim tokens without selling or unblock a stuck channel

6. **Panic Sell**: Sell (Chain A)
   *Use case*: Emergency cleanup if bridging fails after buying

## 🛡️ Risk Management

The bot includes several safety features to protect your capital:

- **Panic Sell**: Automatically sells if bridging fails
- **Nonce Gap Handling**: Detects and resolves stuck transactions
- **Dynamic Fee Buffers**: Adds a safety margin to LayerZero fees
- **Gas Estimation**: Uses optimal gas prices for faster execution
- **Slippage Protection**: Configurable slippage tolerance for swaps

## 🐛 Troubleshooting Common Issues

| Issue | Solution |
|-------|----------|
| "PRIVATE_KEY environment variable not set" | Set the PRIVATE_KEY environment variable as shown above |
| "Failed to connect to RPC" | Check your RPC URL in config.json; ensure your provider is online |
| "Invalid recipe type: 0" | Verify the RecipeType enum is correctly implemented |
| "Gas estimation failed" | Ensure your wallet has enough ETH for gas; try increasing gas limit |
| "Execution failed" | Check OFT address is correct; ensure allowance is set for DEX aggregator |

## 📚 Learning Resources for Beginners

New to Go? Check out these resources:

- [Go Tour](https://go.dev/tour/welcome/1) - Interactive Go tutorial
- [Go by Example](https://gobyexample.com/) - Simple, practical Go examples
- [Go by Example (YouTube)](https://www.youtube.com/watch?v=YS4e4q9oBaU&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6) - Video tutorials

## 📝 Next Steps

1. **Test on Testnets First**: Use Goerli, Sepolia, or BSC Testnet before mainnet
2. **Start Small**: Use small amounts to test the bot's behavior
3. **Monitor Logs**: Watch the logs to understand the bot's decision-making
4. **Customize Strategies**: Modify recipe logic in `strategy/controller.go` to add your own strategies

## 🤝 Contributing

Contributions are welcome! Here's how:

1. Fork the repository
2. Create a new branch (`git checkout -b feature/your-feature`)
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

---

Happy arbitraging! 🚀 Remember to prioritize security and test thoroughly before using real funds.