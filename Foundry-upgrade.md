以下是使用 Foundry 实现多账户管理和合约升级的完整指南：

## 一、生成测试账户

    使用 Anvil 生成带固定助记词的账户
    Bash 复制
    anvil --mnemonic "test test test test test test test test test test test junk"

通过 Cast 生成新账户

    Bash 复制
    # 生成 3 个新账户
    for i in {1..3}; do
      cast wallet new --unsafe-password password | tee account$i.txt
    done

Foundry 的 Anvil 节点默认提供 10 个测试账户（含私钥），也可自定义生成：

    Bash  复制
    # 生成新账户（私钥 + 地址）
    cast wallet new
    # 输出示例：
    # Private Key: 0x2a8718a9d99123b8f225a0400c5874681db596d03b726c5eb90d38726b5e1ca4
    # Address: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266


## 二、基础合约准备

可升级合约 V1 版本

    Solidity 复制
    // SPDX-License-Identifier: MIT
    pragma solidity ^0.8.20;
    
    // 公共存储槽
    contract Storage {
        uint256 public value;
        
        // UUPS 需要保留该插槽
        bytes32[50] private __gap;
    }
    
	
    // 透明代理版本
    contract LogicV1Transparent is Storage {
        function initialize() public {
            value = 100;
        }
        
        function add(uint256 x) public {
            value += x;
        }
    }
    
	
    // UUPS 版本
    import "@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol";
    
    contract LogicV1UUPS is Storage, UUPSUpgradeable {
        function initialize() public initializer {
            value = 100;
        }
        
        function add(uint256 x) public {
            value += x;
        }
        
        // UUPS 必须实现的升级授权
        function _authorizeUpgrade(address) internal override onlyOwner {}
    }


## 三、部署与升级操作

透明代理流程

部署脚本 script/DeployTransparent.s.sol

    Solidity 复制
    
    import {Script} from "forge-std/Script.sol";
    import {ERC1967Proxy} from "openzeppelin-contracts/proxy/ERC1967/ERC1967Proxy.sol";
    
    contract DeployTransparent is Script {
        function run() external {
            address admin = vm.envAddress("ADMIN_ADDRESS");
            
            vm.startBroadcast(admin);
            LogicV1Transparent impl = new LogicV1Transparent();
            ERC1967Proxy proxy = new ERC1967Proxy(
                address(impl),
                abi.encodeWithSelector(LogicV1Transparent.initialize.selector)
            );
            vm.stopBroadcast();
        }
    }
	
执行部署

    Bash 复制
    export ADMIN_ADDRESS=0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
    forge script script/DeployTransparent.s.sol --broadcast --rpc-url http://localhost:8545

UUPS 流程

部署脚本 script/DeployUUPS.s.sol

    Solidity 复制
    contract DeployUUPS is Script {
        function run() external {
            address owner = vm.envAddress("OWNER_ADDRESS");
            
            vm.startBroadcast(owner);
            LogicV1UUPS impl = new LogicV1UUPS();
            ERC1967Proxy proxy = new ERC1967Proxy(
                address(impl),
                abi.encodeWithSelector(LogicV1UUPS.initialize.selector)
            );
            vm.stopBroadcast();
        }
    }

执行部署


    Bash 复制
    export OWNER_ADDRESS=0x70997970C51812dc3A010C7d01b50e0d17dc79C8
    forge script script/DeployUUPS.s.sol --broadcast --rpc-url http://localhost:8545
    

## 四、合约升级操作

准备 V2 合约

    Solidity 复制
    contract LogicV2Transparent is Storage {
        function multiply(uint256 x) public {
            value *= x;
        }
    }
    
    contract LogicV2UUPS is Storage, UUPSUpgradeable {
        function multiply(uint256 x) public {
            value *= x;
        }
        
        function _authorizeUpgrade(address) internal override onlyOwner {}
    }

透明代理升级脚本


    Solidity 复制
    contract UpgradeTransparent is Script {
        function run() external {
            address proxy = vm.envAddress("PROXY_ADDRESS");
            address admin = vm.envAddress("ADMIN_ADDRESS");
            
            vm.startBroadcast(admin);
            LogicV2Transparent v2 = new LogicV2Transparent();
            (bool success,) = proxy.call(abi.encodeWithSignature("upgradeTo(address)", address(v2)));
            require(success, "Upgrade failed");
            vm.stopBroadcast();
        }
    }

UUPS 升级脚本

    Solidity 复制
    contract UpgradeUUPS is Script {
        function run() external {
            address proxy = vm.envAddress("PROXY_ADDRESS");
            address owner = vm.envAddress("OWNER_ADDRESS");
            
            vm.startBroadcast(owner);
            LogicV2UUPS v2 = new LogicV2UUPS();
            LogicV1UUPS(payable(proxy)).upgradeTo(address(v2));
            vm.stopBroadcast();
        }
    }


## 五、验证升级结果

交互测试命令

    Bash复制
    # 检查实现地址
    cast implementation --rpc-url http://localhost:8545 <PROXY_ADDRESS>
    
    # 调用新方法
    cast send <PROXY_ADDRESS> "multiply(uint256)" 2 \
    --private-key $OWNER_PRIVATE_KEY \
    --rpc-url http://localhost:8545
    
自动化测试用例

    Solidity 复制
    function test_Upgrade() public {
        // 初始状态
        assertEq(proxy.value(), 100);
        
        // 执行升级
        vm.prank(admin);
        proxy.upgradeTo(address(v2));
        
        // 调用新功能
        proxy.multiply(2);
        assertEq(proxy.value(), 200);
        
        // 验证旧功能
        proxy.add(50);
        assertEq(proxy.value(), 250);
    }


## 六、安全注意事项

1 权限管理

    Solidity 复制
    // 透明代理需检查 msg.sender 是否为 admin
    modifier onlyAdmin() {
        require(msg.sender == _getAdmin(), "Not authorized");
        _;
    }

    // UUPS 使用 OpenZeppelin 的 Ownable 控制
    import "@openzeppelin/contracts/access/Ownable.sol";
    
    contract LogicV1UUPS is Storage, UUPSUpgradeable, Ownable {}
    

2 存储保护
保持存储变量的顺序不变
新增变量始终添加到末尾
使用 __gap 保留存储空间


3 升级验证

    Bash 复制
    # 检查存储插槽
    cast storage <PROXY_ADDRESS> 0 --rpc-url localhost

