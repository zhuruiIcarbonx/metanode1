# metanode1

https://leetcode.cn/problems/single-number/
https://leetcode.cn/problems/palindrome-number/
https://leetcode.cn/problems/valid-parentheses/
https://leetcode.cn/problems/longest-common-prefix/
https://leetcode.cn/problems/plus-one/
https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
https://leetcode.cn/problems/merge-intervals/
https://leetcode.cn/problems/two-sum/
112233
https://github.com/mao888/golang-guide/blob/main/golang/go-Interview/GOALNG_INTERVIEW_COLLECTION.md
https://github.com/DCreek03/web3.0-learning/tree/main/go_work/task4
https://github.com/MetaNodeAcademy/LearningRoadmap/blob/main/%E8%BF%9B%E9%98%B61-DApp%E5%90%8E%E7%AB%AF%E5%AE%9E%E6%88%98.md

https://go-zero.dev/docs/tutorials
https://space.bilibili.com/389552232/upload/video

//视频gozero例子
https://github.com/Mikaelemmmm/

//wakuang
https://ncnlo95gua0y.feishu.cn/wiki/FwZQwnetZiwIwfkdR3Eckjvanod

//english
https://www.oxfordlearnersdictionaries.com/wordlist/american_english/oxford3000/


1、完成补充智能合约内容的学习
参考资料：https://github.com/AmazingAng/WTF-Solidity?tab=readme-ov-file
从第一讲学习到S17跨服重入攻击

问题：
1、contract C is A {
    constructor(uint _c) A(_c * _c) {}
}

A(_c * _c)

2、CREATE2是否会重复创建


Rust:
1、https://github.com/rust-boom/rust-boom
2、https://github.com/dodolalorc/BlockChain
3、https://rustmagazine.github.io/rust_magazine_2021/chapter_5/six-years-of-rust.html

------------------------
    FT资料: https://github.com/locey/NFT_hello_world/blob/main/01_how_to_build_NFT.md
	别人作业：https://github.com/caryxiao/NFT-Auction
	
	下周学习内容
	1、完成Meme代币系列
	2、MetaNodeStake
	参考链接：https://github.com/MetaNodeAcademy/LearningRoadmap/blob/main/%E8%BF%9B%E9%98%B62-%E5%90%88%E7%BA%A6%E5%BC%80%E5%8F%91%E5%AE%9E%E6%88%98.m
	完成任务1，2之后提交gitgub链接到discord
	
	ethers.js的教程：https://github.com/WTFAcademy/WTF-Ethers

------------------------

大作业：实现一个 NFT 拍卖市场
任务目标
使用 Hardhat 框架开发一个 NFT 拍卖市场。
使用 Chainlink 的 feedData 预言机功能，计算 ERC20 和以太坊到美元的价格。
使用 UUPS/透明代理模式实现合约升级。
使用类似于 Uniswap V2 的工厂模式管理每场拍卖。


任务步骤
项目初始化
使用 Hardhat 初始化项目：
npx hardhat init
安装必要的依赖：
  npm install @openzeppelin/contracts @chainlink/contracts @nomiclabs/hardhat-ethers hardhat-deploy
实现 NFT 拍卖市场
NFT 合约：
使用 ERC721 标准实现一个 NFT 合约。
支持 NFT 的铸造和转移。
拍卖合约：
实现一个拍卖合约，支持以下功能：
创建拍卖：允许用户将 NFT 上架拍卖。
出价：允许用户以 ERC20 或以太坊出价。
结束拍卖：拍卖结束后，NFT 转移给出价最高者，资金转移给卖家。
工厂模式：
使用类似于 Uniswap V2 的工厂模式，管理每场拍卖。
工厂合约负责创建和管理拍卖合约实例。
集成 Chainlink 预言机
价格计算：
使用 Chainlink 的 feedData 预言机，获取 ERC20 和以太坊到美元的价格。
在拍卖合约中，将出价金额转换为美元，方便用户比较。
跨链拍卖：
使用 Chainlink 的 CCIP 功能，实现 NFT 跨链拍卖。
允许用户在不同链上参与拍卖。
合约升级
UUPS/透明代理：
使用 UUPS 或透明代理模式实现合约升级。
确保拍卖合约和工厂合约可以安全升级。
测试与部署
测试：
编写单元测试和集成测试，覆盖所有功能。
部署：
使用 Hardhat 部署脚本，将合约部署到测试网（如 Goerli 或 Sepolia）






Go-Ethereum 核心功能与架构设计研究作业
作业目的
通过本作业深入理解以太坊参考实现 Go-Ethereum（Geth）的设计哲学，掌握区块链核心组件的实现原理

任务分解
理论分析（40%）
1 阐述Geth在以太坊生态中的定位
2 解析核心模块交互关系：
区块链同步协议（eth/62,eth/63）
交易池管理与Gas机制
EVM执行环境构建
共识算法实现（Ethash/POS）


架构设计（30%）
绘制分层架构图（需包含以下层级）：
[P2P网络层]->[区块链协议层]->[状态存储层]->[EVM执行层]
说明各层关键模块：
les（轻节点协议）
trie（默克尔树实现）
core/types（区块数据结构）

实践验证（30%）
1 编译并运行Geth节点
make geth
./build/bin/geth --dev --http
2 通过控制台验证功能：
eth.blockNumber // 查看区块高度
miner.start()   // 启动挖矿


作业要求
1 研究报告需包含：
功能架构图交易生命周期流程图
账户状态存储模型

2 实践报告需包含：
私有链搭建过程
智能合约部署截图
区块浏览器查询结果

参考资料
官方架构文档：https://geth.ethereum.org/docs
源码精读路线：
/go-ethereum
├── core      // 区块链核心逻辑
├── miner     // 挖矿模块
├── eth       // 以太坊协议实现
└── internal  // 底层工具包

评分标准：架构完整性40%、实现深度30%、实践完成度30%

该作业通过源码分析、模块调试和系统设计三个维度，培养区块链底层开发能力。需要特别注意P2P网络层的kademlia协议实现和状态数据库的MPT树结构













// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test} from "forge-std/Test.sol";
import {Counter} from "../src/Counter.sol";
import {MetaNodeStake} from "../src/MetaNodeStake.sol";
import {MetaNodeERC20} from "../src/MetaNodeERC20.sol";
import {StakeERC20} from "../src/StakeERC20.sol";
import {console} from "forge-std/console.sol";

contract MetaNodeStakeTest is Test {
    MetaNodeStake public stake;
    StakeERC20 public stakeERC20;
    MetaNodeERC20 public metaNodeERC20;

    function setUp() public {
        stakeERC20 = new StakeERC20();
        metaNodeERC20 = new MetaNodeERC20();
        stake = new MetaNodeStake();
        stake.initialize(metaNodeERC20,10,10000,3_000_000_000_000_000_000); 

    }

    function test_addPool() public {

        address _stTokenAddress = address(0);
        uint256 _poolWeight = 10000;
        uint256 _minDepositAmount = 100;
        uint256 _unstakeLockedBlocks = 100;
        stake.addPool(_stTokenAddress,_poolWeight,_minDepositAmount,_unstakeLockedBlocks);

        (
            address stTokenAddress, //质押代币地址
            uint256 stTokenAmount, //池中总质押代币量
            uint256 poolWeight, //质押池的权重，影响奖励分配
            uint256 lastRewardBlock,//最后一次计算奖励的区块
            uint256 accMetaNodePerST,//每个质押代币积累的奖励代币数量 ----这个参数的计算和使用有问题，该方式应该不适用实际业务
            uint256 minDepositAmount,//最小质押金额
            uint256 unstakeLockedBlocks //解除质押的锁
        )  = stake.pools(0);

        console.log("stTokenAddress-----",stTokenAddress);
        console.log("stTokenAmount-----",stTokenAmount);
        console.log("poolWeight-----",poolWeight);
        console.log("lastRewardBlock-----",lastRewardBlock);
        console.log("accMetaNodePerST-----",accMetaNodePerST);
        console.log("minDepositAmount-----",minDepositAmount);
        console.log("unstakeLockedBlocks-----",unstakeLockedBlocks);
        

        
    }

    
}

