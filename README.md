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



------------------------

任务 1：区块链读写 任务目标
使用 Sepolia 测试网络实现基础的区块链交互，包括查询区块和发送交易。
 具体任务
1 环境搭建
安装必要的开发工具，如 Go 语言环境、 go-ethereum 库。
注册 Infura 账户，获取 Sepolia 测试网络的 API Key。
2 查询区块
编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
实现查询指定区块号的区块信息，包括区块的哈希、时间戳、交易数量等。
输出查询结果到控制台。


3 发送交易
准备一个 Sepolia 测试网络的以太坊账户，并获取其私钥。
编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
构造一笔简单的以太币转账交易，指定发送方、接收方和转账金额。
对交易进行签名，并将签名后的交易发送到网络。
输出交易的哈希值。

任务 2：合约代码生成 任务目标
使用 abigen 工具自动生成 Go 绑定代码，用于与 Sepolia 测试网络上的智能合约进行交互。
 具体任务
1 编写智能合约
使用 Solidity 编写一个简单的智能合约，例如一个计数器合约。
编译智能合约，生成 ABI 和字节码文件。
2 使用 abigen 生成 Go 绑定代码
安装 abigen 工具。
使用 abigen 工具根据 ABI 和字节码文件生成 Go 绑定代码。
3 使用生成的 Go 绑定代码与合约交互
编写 Go 代码，使用生成的 Go 绑定代码连接到 Sepolia 测试网络上的智能合约。
调用合约的方法，例如增加计数器的值。
输出调用结果。




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

