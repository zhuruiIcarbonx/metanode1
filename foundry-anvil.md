## 一、启动本地节点
        启动 Anvil（默认端口 8545）
        Bash 复制
        anvil 

# 保持终端运行，新开终端继续后续操作

    带参数的启动方式（示例）
        Bash 复制
        anvil \
        --port 8888 \          # 自定义端口
        --block-time 5 \       # 每5秒出一个块
        --mnemonic "test" \    # 固定助记词
        --balance 1000         # 每个测试账户初始余额(ETH)

## 二、部署合约到本地节点

编写部署脚本 script/Counter.s.sol

    Solidity 复制
    // SPDX-License-Identifier: MIT
    pragma solidity ^0.8.20;
    
    import {Script} from "forge-std/Script.sol";
    import {Counter} from "../src/Counter.sol";
    
    contract DeployCounter is Script {
        function run() external {
            vm.startBroadcast();
            new Counter();
            vm.stopBroadcast();
        }
    }

执行部署（使用Anvil第一个测试账户）

    Bash复制
    forge script script/Counter.s.sol:DeployCounter \
    --rpc-url http://localhost:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --broadcast


## 三、测试合约交互

方式1：使用 Cast 命令行工具

查询合约状态

    Bash复制
    # 替换 <CONTRACT_ADDRESS> 为实际部署地址
    cast call <CONTRACT_ADDRESS> "count()(uint256)" \
    --rpc-url http://localhost:8545

发送交易调用

    Bash 复制
    cast send <CONTRACT_ADDRESS> "increment()" \
    --rpc-url http://localhost:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

方式2：使用 Foundry 测试框架

创建集成测试文件 test/Integration.t.sol

    Solidity 复制
    // SPDX-License-Identifier: MIT
    pragma solidity ^0.8.20;
    
    import {Test, console} from "forge-std/Test.sol";
    import {Counter} from "../src/Counter.sol";
    
    contract IntegrationTest is Test {
        Counter public counter;
        address user = vm.addr(1); // 测试用户
    
        function setUp() public {
            vm.startPrank(user);
            counter = new Counter();
            vm.stopPrank();
        }
    
        function test_Integration() public {
            vm.startPrank(user);
            counter.increment();
            assertEq(counter.count(), 1);
            
            counter.decrement();
            assertEq(counter.count(), 0);
            vm.stopPrank();
        }
    }
	

运行测试

    Bash 复制
    forge test --match-contract IntegrationTest -vvv

## 四、查看节点信息

查看当前区块

    Bash 复制
    cast block latest --rpc-url http://localhost:8545

查看账户余额

    Bash 复制
    cast balance 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
    --rpc-url http://localhost:8545

查看交易详情

    Bash 复制
    cast tx <TX_HASH> --rpc-url http://localhost:8545

## 五、常用技巧

快速重置环境：

    Bash 复制
    # 关闭原有anvil进程
    pkill anvil

# 重新启动并部署
    anvil & 
    forge script ... --broadcast

调试交易：

Bash复制
cast run <TX_HASH> --debug \
--rpc-url http://localhost:8545

模拟时间流逝：

    Solidity 复制
    // 在测试中向前推进时间
    vm.warp(block.timestamp + 86400); // 快进1天
    

## 六、注意事项


    Anvil 默认创建10个测试账户，每个账户有10000 ETH
    本地节点重启后合约需要重新部署
    使用 vm.startPrank(address) 模拟用户操作
    通过 --gas-limit 可调整交易 Gas 限制
    Anvil 日志会显示所有交易详情（在启动终端查看）
