
## 一、环境准备
1、安装 Foundry

    curl -L https://foundry.paradigm.xyz | bash
    foundryup


2、创建新项目

        forge init my_project
        cd my_project
    

## 二、编写智能合约

在 src/ 目录创建合约文件 Counter.sol


    // SPDX-License-Identifier: MIT
    pragma solidity ^0.8.20;
    
    contract Counter {
        uint256 public count;
    
        function increment() public {
            count += 1;
        }
    
        function decrement() public {
            require(count > 0, "Counter: cannot decrement below zero");
            count -= 1;
        }
    }

## 三、编写测试用例

在 test/ 目录创建测试文件 Counter.t.sol
Solidity


    // SPDX-License-Identifier: MIT
    pragma solidity ^0.8.20;
    
    import {Test} from "forge-std/Test.sol";
    import {Counter} from "../src/Counter.sol";
    
    contract CounterTest is Test {
        Counter public counter;
    
        function setUp() public {
            counter = new Counter();
        }
    
        function test_Increment() public {
            counter.increment();
            assertEq(counter.count(), 1);
        }
    
        function test_DecrementRevert() public {
            vm.expectRevert("Counter: cannot decrement below zero");
            counter.decrement();
        }
    }

运行测试

`forge test -vvv  # -vvv 显示详细日志`



## 四、部署合约

1、创建部署脚本 script/Counter.s.sol

    // SPDX-License-Identifier: MIT
    pragma solidity ^0.8.20;
    
    import {Script} from "forge-std/Script.sol";
    import {Counter} from "../src/Counter.sol";
    
    contract DeployCounter is Script {
        function run() external returns (Counter) {
            vm.startBroadcast();
            Counter counter = new Counter();
            vm.stopBroadcast();
            return counter;
        }
    }

2、部署到本地节点（需先启动 Anvil）

    `anvil &  # 启动本地节点
    
    forge script script/Counter.s.sol:DeployCounter --rpc-url http://localhost:8545 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --broadcast
`

部署到测试网（示例：Sepolia）

    forge script script/Counter.s.sol:DeployCounter \
    --rpc-url $SEPOLIA_RPC_URL \
    --private-key $PRIVATE_KEY \
    --broadcast \
    --verify -e etherscan-api-key

## 五、常用命令速查

    功能	命令
    安装依赖	forge install owner/repo
    编译合约	forge build
    查看测试覆盖率	forge coverage
    生成 Gas 报告	forge test --gas-report
    交互式调试	forge test --debug <test_function>

## 六、注意事项

私钥管理建议使用 .env 文件：

    # .env 文件示例
    PRIVATE_KEY=your_private_key
    SEPOLIA_RPC_URL=https://sepolia.infura.io/v3/your-project-id

配置 foundry.toml 优化编译：

    Toml
    
    [profile.default]
    src = "src"
    out = "out"
    libs = ["lib"]
    solc = "0.8.20"
    optimizer = true
    optimizer_runs = 200

本地开发推荐组合：
    
    Bash
    
    anvil &          # 启动本地节点
    forge test       # 运行测试
    cast send ...    # 发送交易
