ts手册
https://www.typescriptlang.org/docs/handbook/2/generics.html

安装
npm install -g typescript

检查错误
tsc hello.ts





项目：
https://github.com/Cyfrin/foundry-simple-storage-cu/blob/main/test/SimpleStorageTest.t.sol

https://github.com/Cyfrin/foundry-full-course-cu?tab=readme-ov-file#intro

https://github.com/adshao/publications/blob/master/uniswap/dive-into-uniswap-v3-whitepaper/README.md

https://www.typescriptlang.org/docs/handbook/typescript-in-5-minutes-oop.html


Foundry学习

1、入门非常简单：

	安装foundryup： curl -L https://foundry.paradigm.xyz | bash

	接下来，运行： foundryup 。
	它将自动安装最新版本的预编译二进制文件：forge和cast anvil chisel
	foundryup

	完毕！

2、forge

	Forge 可帮助您构建、测试、模糊测试、调试和部署 Solidity 合约。了解 Forge 的最好方法就是简单尝试一下（不到 30 秒！）。
	首先，让我们初始化一个新的counter示例存储库：
	 
	forge init counter

	接下来cd进入counter并构建：

	cd counter

	forge build


	让我们测试一下我们的合约：

	forge test

	最后，让我们运行部署脚本：

	forge script script/Counter.s.sol

	运行forge --help以探索可用子命令的完整列表及其用法。
	更多文档可以在Foundry Docs 的forge部分找到



3、cast


	Cast 是一把瑞士军刀，用于从命令行与以太坊应用程序交互。以下是您可以执行的操作的一些示例：

	检查以太坊主网上的最新区块：
	cast block-number --rpc-url https://eth.merkle.io

	检查以太币余额vitalik.eth
	cast balance vitalik.eth --ether --rpc-url https://eth.merkle.io

	重放并追踪交易
	cast run 0x9c32042f5e997e27e67f82583839548eb19dc78c4769ad6218657c17f2a5ed31 --rpc-url https://eth.merkle.io

  





4、Anvil
	Anvil 是一个快速的本地以太坊开发节点。
	让我们在最新区块处分叉以太坊主网：

	anvil --fork-url https://eth.merkle.io

	您可以cast对您的anvil实例使用相同的子命令：
	cast block-number

5、Chisel

	Chisel 是一种快速、实用且详细的 Solidity REPL。
	要使用 Chisel，只需输入chisel。

	chisel

	从这里开始编写 Solidity 代码吧！Chisel 会对每个输入提供详细的反馈。
	创建一个变量a并查询它：
	➜ uint256 a = 123;
	➜ a
	Type: uint256
	├ Hex: 0x7b
	├ Hex (full word): 0x000000000000000000000000000000000000000000000000000000000000007b
	└ Decimal: 123


6、配置

foundry.tomlFoundry 具有高度可配置性，您可以根据自己的需求进行定制。配置通过位于项目根目录或任何父目录中的名为 的文件进行管理。有关配置选项的完整列表，请参阅config 包文档。

配置文件和命名空间

配置可以组织成配置文件，这些配置文件可以任意命名，以实现灵活性。
默认配置文件名为default。请参阅“默认配置文件”部分了解更多信息。
要选择不同的配置文件，请设置FOUNDRY_PROFILE环境变量。
使用以 为前缀的环境变量覆盖特定设置FOUNDRY_（例如FOUNDRY_SRC）。
