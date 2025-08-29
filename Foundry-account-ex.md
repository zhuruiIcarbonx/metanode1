在 Foundry 中配置测试账号并获取其地址/私钥，可通过以下步骤实现：

## 1. 生成测试账号
### 方法 1：使用 Anvil 自动生成功能
启动 Anvil 节点时，系统默认生成 10 个测试账户（含私钥）：

    Bash复制
    forge anvil
    
    # 输出示例：
    # 管理员账户: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
    # 私钥: 0xac0974bec39a17e36baa7a2e0d33c5eeaad7933cb1969b4393aee650ef9a756f

### 方法 2：手动创建新账户
使用 cast wallet new 生成新账户：

    Bash复制
    cast wallet new
    # 输出：
    # Private Key: 0x2a8718a9d99123b8f225a0400c5874681db596d03b726c5eb90d38726b5e1ca4
    # Address: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266

## 2. 配置测试账号

### 方式 1：通过环境变量存储私钥
在 .env 文件中定义私钥（避免硬编码）：

    Bash复制
    # .env
    PRIVATE_KEY=0xac0974bec39a17e36baa7a2e0d33c5eeaad7933cb1969b4393aee650ef9a756f
    在脚本中读取：
    
    Solidity复制
    // script/Deploy.s.sol
    uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");

### 方式 2：在 foundry.toml 中配置

    可预定义多个账户别名：
    
    Toml复制
    [env-aliases]
    deployer = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
    tester = "0xAbc...123"

## 3. 部署时使用测试账号

示例：使用 JavaScript 脚本部署

    Javascript复制
    // script/deploy.js
    const hre = require("hardhat");
    
    async function main() {
        const [deployer] = await hre.ethers.getSigners(); // 获取默认账户
        console.log("Deploying contracts with account:", deployer.address);
    
        const MyContract = await hre.ethers.getContractFactory("MyContract");
        const contract = await MyContract.deploy();
        await contract.waitForDeployment();
        console.log("Contract deployed to:", await contract.getAddress());
    }

main();
指定私钥部署
通过命令行传递私钥：

    Bash复制
    forge script script/deploy.js \
      --rpc-url http://127.0.0.1:8545 \
      --private-key 0xac0974bec39a17e36baa7a2e0d33c5eeaad7933cb1969b4393aee650ef9a756f \
      --broadcast

## 4. 测试时获取测试账号

在 Solidity 测试用例中
使用 forge-std 的 cheats 模块：

    Solidity复制
    // test/MyContract.t.sol
    pragma solidity ^0.8.23;
    
    import "forge-std/Test.sol";
    import "../src/MyContract.sol";
    
    contract MyContractTest is Test {
        MyContract public contract;
        address user = makeAddr("user"); // 生成测试账号
    
        function setUp() public {
            contract = new MyContract();
        }
    
        function testTransfer() public {
            vm.prank(user); // 模拟 user 发起交易
            contract.transfer(address(1), 100);
        }
    }

批量生成测试账号

    Solidity复制
    address[] memory users = new address[](10);
    for (uint256 i = 0; i < 10; i++) {
        users[i] = makeAddr(string.concat("user", Strings.toString(i))); // 生成 user0 ~ user9
    }

## 5. Fork 模式下使用主网账户

若需在 Fork 环境中使用特定主网账户：

    Bash复制
    # 启动 Fork 节点
    forge anvil --fork-url https://eth-mainnet.g.alchemy.com/v2/your-api-key
    
    # 使用主网账户私钥部署
    forge script script/deploy.js \
      --rpc-url http://127.0.0.1:8545 \
      --private-key 0x主网账户私钥 \
      --broadcast

## 6. 注意事项

    私钥安全：避免在代码中硬编码私钥，优先使用环境变量。
    测试账号管理：
    使用 makeAddr(string memory label) 生成可复现的测试账号（相同 label 生成相同地址）。
    使用 vm.deal(address, uint256) 快速充值测试账户 ETH。
通过以上步骤，您可以在 Foundry 中灵活配置和使用测试账号，实现合约部署与测试的全流程管理。
