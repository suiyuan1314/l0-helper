# LayerZero V2 模块化套利机器人

一个基于LayerZero V2构建的模块化高频套利（HFA）机器人框架，采用"乐高积木"架构设计。该机器人可以动态组合标准化模块，在EVM兼容链上形成各种套利策略。

## 📋 前置条件

开始之前，请确保您已准备好以下内容：

1. **Go 1.25.4 或更新版本** - [下载 Go](https://go.dev/dl/)
2. **代码编辑器** - VS Code（推荐）、GoLand 或任何文本编辑器
3. **Git** - 用于克隆仓库
4. **以太坊钱包** - 带有测试网ETH/稳定币用于测试
5. **RPC端点** - 用于您要套利的链（Alchemy、Infura 或公共RPC）

## 📂 项目结构

机器人采用**模块化"乐高积木"架构**，每个组件（引擎）独立工作，可以组合创建不同策略：

```
├── abis/               # LayerZero V2 合约ABI
├── bridge/             # LayerZero 桥接操作（发送/接收）
├── config/             # 配置管理
├── contracts/          # 智能合约交互
├── execution/          # 交易管理，支持打包
├── guard/              # 风险验证和nonce间隙处理
├── liquidity/          # DEX聚合器交互，获取最优交换路径
├── monitoring/         # 日志和监控
├── risk/               # 高级风险管理
├── strategy/           # 编排所有六种套利策略
├── config.json         # 机器人配置文件
├── go.mod              # Go依赖
└── main.go             # 测试入口点
```

## 🚀 快速开始

### 1. 克隆仓库

打开终端并克隆此仓库：

```bash
git clone https://github.com/your-username/l0-modular-arb-bot.git
cd l0-modular-arb-bot
```

### 2. 安装依赖

Go模块自动管理依赖。运行以下命令下载所有必需的包：

```bash
go mod tidy
```

### 3. 配置机器人

编辑 `config.json` 以设置您的链和偏好。以下是配置文件的详细说明：

```json
{
  "chains": [
    {
      "name": "Ethereum",
      "eid": 30101,              // LayerZero 链ID
      "chainId": 1,             // 以太坊主网链ID（测试网Goerli使用5）
      "rpcUrl": "https://eth-mainnet.g.alchemy.com/v2/YOUR_KEY",
      "wsUrl": "wss://eth-mainnet.g.alchemy.com/v2/YOUR_KEY",
      "contracts": {
        "endpoint": "0x1a44076050125825900e736c501f859c50fE728c",  // LayerZero EndpointV2
        "endpointView": "",    // EndpointV2View 合约
        "oft": "0x...",        // OFT（跨链代币）地址
        "aggregatorSpender": "0x1111111254fb6c44bac0bed2854e76f90643097d"  // DEX聚合器授权地址
      },
      "aggregator": {
        "apiUrl": "https://api.1inch.dev/swap/v6.0/1/swap",  // DEX聚合器API
        "apiKey": "YOUR_1INCH_KEY"
      },
      "bundling": {
        "enabled": false,       // 启用Flashbots/Jito打包
        "relayUrl": "https://relay.flashbots.net"
      }
    }
  ],
  "bot": {
    "privateKey": "ENV_VAR",    // 使用环境变量存储私钥
    "maxGapThreshold": 3,       // 最大可解决的nonce间隙
    "feeBufferMultiplier": 1.1  // LayerZero费用安全缓冲倍数
  }
}
```

**初学者快速提示**：
- 先使用测试网（如以太坊Goerli、BSC测试网）
- 从水龙头获取免费测试网代币，如 [Goerli Faucet](https://goerlifaucet.com/)
- 如果没有Alchemy/Infura密钥，可以使用公共RPC

### 4. 设置环境变量

永远不要硬编码私钥！请使用环境变量：

**Mac/Linux**：
```bash
export PRIVATE_KEY="你的私钥"
```

**Windows（命令提示符）**：
```cmd
set PRIVATE_KEY=你的私钥
```

**Windows（PowerShell）**：
```powershell
$env:PRIVATE_KEY="你的私钥"
```

## 🔧 测试机器人

运行测试脚本以验证一切正常工作：

```bash
go run main.go
```

您应该会看到类似输出：
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

## 🍲 套利策略

机器人支持6种套利策略（"食谱"）：

1. **标准套利**：购买（链A）→ 桥接 → 出售（链B）
   *使用场景*：资金为ETH/稳定币；检测到跨链套利机会

2. **库存套利**：桥接（链A）→ 出售（链B）
   *使用场景*：机器人已在源链持有OFT；目标链出现价格差距

3. **再平衡**：桥接（链A）→ 接收（链B）
   *使用场景*：将库存转移到交易量更高或未来机会更好的链

4. **"终结者"**：接收（链B）→ 出售（链B）
   *使用场景*：完成进行中的套利或"捕获"机器人重启后的数据包

5. **"信使"**：纯接收（链B）
   *使用场景*：仅领取代币而不出售，或解锁卡住的通道

6. **恐慌出售**：出售（链A）
   *使用场景*：如果在购买后桥接失败，紧急清理源链库存

## 🛡️ 风险管理

机器人包含多种安全功能，保护您的资金：

- **恐慌出售**：如果桥接失败，自动出售
- **Nonce间隙处理**：检测并解决卡住的交易
- **动态费用缓冲**：为LayerZero费用添加安全边际
- **Gas估算**：使用最优Gas价格加快执行速度
- **滑点保护**：可配置的交换滑点容忍度

## 🐛 常见问题故障排除

| 问题 | 解决方案 |
|-------|----------|
| "PRIVATE_KEY environment variable not set" | 按照上述方法设置PRIVATE_KEY环境变量 |
| "Failed to connect to RPC" | 检查config.json中的RPC URL；确保您的提供商在线 |
| "Invalid recipe type: 0" | 验证RecipeType枚举是否正确实现 |
| "Gas estimation failed" | 确保钱包有足够的ETH支付Gas；尝试增加Gas限制 |
| "Execution failed" | 检查OFT地址是否正确；确保已为DEX聚合器设置授权 |

## 📚 初学者学习资源

不熟悉Go语言？查看以下资源：

- [Go语言之旅](https://go.dev/tour/welcome/1) - 交互式Go教程
- [Go示例](https://gobyexample.com/) - 简单实用的Go示例
- [Go编程入门（B站）](https://www.bilibili.com/video/BV1gf4y1r79E) - 适合中文用户的视频教程

## 📝 后续步骤

1. **先在测试网上测试**：在主网之前，先使用Goerli、Sepolia或BSC测试网
2. **从小金额开始**：使用小额资金测试机器人行为
3. **监控日志**：观察日志以了解机器人的决策过程
4. **自定义策略**：修改`strategy/controller.go`中的策略逻辑，添加您自己的策略

## 🤝 贡献

欢迎贡献！步骤如下：

1. Fork 本仓库
2. 创建新分支（`git checkout -b feature/your-feature`）
3. 进行修改
4. 彻底测试
5. 提交拉取请求

## 📄 许可证

本项目采用MIT许可证 - 有关详情，请参阅LICENSE文件。

---

套利愉快！🚀 请记住，在使用真实资金之前，务必优先考虑安全并彻底测试。