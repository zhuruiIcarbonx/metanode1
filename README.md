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

