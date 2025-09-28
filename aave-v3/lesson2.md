入门：设置你的 DeFi Aave V3 环境
欢迎来到 DeFi Aave V3 课程！本指南将指导您访问所有必要的课程资源并配置您的开发环境。成功设置环境对于完成实践练习和充分理解 Aave V3 协议至关重要。

访问课程资料：核心 GitHub 存储库
所有主要课程材料，包括练习、笔记和解决方案，都托管在私人 GitHub 存储库中：Cyfrin/defi-aave-v3。

首先，您需要将此存储库克隆到本地计算机。

导航到存储库的 GitHub 页面（例如https://github.com/Cyfrin/defi-aave-v3）。

单击绿色的“代码”按钮。

复制提供的 SSH URL（例如git@github.com:Cyfrin/defi-aave-v3.git）。

打开终端并运行以下命令：

git克隆git@github.com:Cyfrin/defi-aave-v3.git
克隆过程完成后，导航到新创建的目录：

cd defi-aave-v3
设置您的 Foundry 开发环境
课程练习是使用 Foundry 框架编写的，位于foundry/您刚刚克隆的存储库的子目录中。

导航到 Foundry 目录：

CD代工厂/
安装依赖项并构建合约：
该项目使用 git 子模块和定义来foundry.toml管理依赖项，例如 Forge-STD、OpenZeppelin 合约和 Aave V3 接口。要安装这些并编译合约，请运行：

锻造建造
使用.env文件配置主网分叉：
这些练习旨在针对以太坊主网的分叉执行。这需要一个主网 RPC（远程过程调用）URL。

在该foundry/目录中，您将找到一个名为 的示例环境文件.env.sample。它包含一个占位符：FORK_URL=。

您需要.env通过复制示例来创建自己的文件：

cp .env.sample . env
使用文本编辑器打开新创建的.env文件，并将占位符替换为您的以太坊主网 RPC URL。例如：

FORK_URL = https://eth-mainnet.g.alchemy.com/v2/YOUR_ALCHEMY_API_KEY​
您可以从Alchemy等提供商处获取免费的 RPC URL 。注册一个帐户并创建一个新的以太坊主网应用程序即可获取您的 URL。

了解 Aave V3：官方合约存储库
要深入了解 Aave V3 协议的实际智能合约，您可以参考官方的 Aave V3 源代码。源代码位于：github.com/aave-dao/aave-v3-origin。

对于本课程，我们将特别关注与v3.3.0该存储库的标签相对应的代码库。

补充学习：课程笔记、图表和 Python 模拟
您克隆的仓库defi-aave-v3还包含一个notes/文件夹。此目录包含有价值的补充材料：

图表：用于解释复杂概念的视觉辅助工具（例如，apr-apy.png说明 APR 和 APY 之间的区别）。

Python 模拟： Jupyter 笔记本（例如binomial_expansion.ipynb），用于通过 Python 代码探索某些机制。

要运行这些 Python 模拟，您需要安装 Jupyter Notebook。文件夹notes/course-setup.md中的文件notes/提供了 Jupyter 的安装链接和说明。
注意：安装 Jupyter 是可选的。核心智能合约练习不需要安装 Jupyter，讲师将演示模拟过程。

使用 Tenderly 掌握交易分析
Tenderly ( tenderly.co ) 是一个强大的平台，用于调试和分析以太坊主网上的交易。它是理解现实世界中 Aave V3 交互方式的宝贵工具。

温柔地使用：

在 Tenderly 上注册一个免费帐户。

导航到您的 Tenderly 仪表板。

将交易哈希（例如0xc1120138b3aa3dc6a49ef7e84eccd17530c273e2442f83e47025d819d9a700743，这是向 Aave V3 提供的 DAI）粘贴到搜索栏中。

需要分析的关键 Tenderly 特征：

摘要：提供交易的概述。

合同：列出交易期间交互的所有合同。

事件：显示所有发出的事件。

状态：详细说明发生的状态变化。

调试器/跟踪：这个功能特别有用。它显示：

合约调用和内部函数调用的顺序。

传递给每个函数的参数。

能够单击跟踪中的一行并查看已执行的确切 Solidity 代码。

通过检查 Tenderly 上的真实交易，您可以深入了解要调用哪些函数以及要为练习传递哪些参数。

安装 Foundry：您的智能合约工具包
Foundry 是本课程所有练习使用的开发框架。如果您尚未安装，请按照以下步骤操作：

官方 Foundry 存储库是github.com/foundry-rs/foundry。

安装（Linux/macOS）：

README.mdFoundry GitHub 页面上的 会将您引导至getfoundry.sh。

安装命令通常为：

curl -L https://foundry.paradigm.xyz | bash
在您的终端中执行此命令。

脚本完成后，按照屏幕上的说明进行操作，通常包括运行foundryup以完成安装并将 Foundry 添加到您的 PATH。

铸造厂
导航和完成课程练习
README.md位于存储库根目录中的主文件defi-aave-v3作为您的课程大纲和目录。

打开README.md文件。

滚动到您正在处理的相关部分和主题（例如，“合同架构”->“供应”）。

在每个主题下，你会发现一个“练习”子部分，其中包含三个重要链接：

练习（Markdown 文件）：链接到一个 Markdown 文件（例如foundry/exercises/Supply.md），其中包含：

练习的详细说明。

要完成的具体任务（例如，“任务 1 - 向 Aave V3 池提供代币”）。

提示来指导您。

运行解决方案的命令forge test（例如forge test --fork-url $FORK_URL --fork-block-number $FORK_BLOCK_NUM --match-path test/Supply.test.sol -vvv）。

启动代码（Solidity 文件）：链接到foundry/src/exercises/Supply.sol您将编写代码的 Solidity 文件（例如 ）。它包含函数存根和概述任务的注释。

答案（Solidity 文件）：此链接指向本练习的完整正确答案（例如foundry/src/solutions/Supply.sol）。在查看答案之前，请先尝试自己解决！

优化您的工作流程：铸造测试技巧和故障排除
README.md目录中的文件包含defi-aave-v3/foundry/专门用于 Foundry 练习的有用命令和注释。

关键命令和提示：

设置.env：

cp .env.sample . env 
（然后.env按照FORK_URL前面描述的方式进行编辑。）

制定你的练习：

锻造建造
构建提供的解决方案：
要构建解决方案代码而不是练习代码，请使用solution配置文件：

FOUNDRY_PROFILE=解决方案锻造构建
优化测试速度FORK_BLOCK_NUM：
在分叉模式下运行测试时，Foundry 会通过您的 RPC URL 从主网获取区块链状态。为了加快速度，您可以将状态缓存在特定的区块编号处。

获取最新的区块号并将其存储在环境变量中：

导出FORK_BLOCK_NUM=$(cast 区块号 --rpc-url $FORK_URL )
# 或者为了跨会话持久化，将其添加到 shell 的 rc 文件（例如，.bashrc、.zshrc）
# 或者每次为项目打开新终端时运行它
使用缓存块测试练习：

伪造测试--fork-url $FORK_URL --fork-block-number $FORK_BLOCK_NUM --match-path测试/Supply.test.sol -vvv
首次运行将从 获取并缓存状态。后续使用相同状态的$FORK_BLOCK_NUM运行将显著加快速度，因为它们使用了缓存状态。 $FORK_BLOCK_NUM

使用缓存块测试解决方案：

FOUNDRY_PROFILE=解决方案锻造测试--fork-url $FORK_URL --fork-block-number $FORK_BLOCK_NUM --match-path测试/Supply.test.sol -vvv
排除测试问题：
如果您遇到意外的测试失败或问题，尤其是在构建练习和解决方案之间切换时，构建工件可能会导致问题。

运行forge clean以删除构建工件和缓存：

锻造清洁
然后，再次尝试构建和测试。

最新状态的替代方案：如果forge clean问题仍然无法解决，或者您始终希望针对最新的主网状态进行测试（这会牺牲速度），请--fork-block-number $FORK_BLOCK_NUM从测试命令中删除该部分。这将强制 Foundry 每次都获取最新状态。

从实际案例中学习：主网交易分析
主目录README.md（位于 的根目录defi-aave-v3）还包含一个“交易”部分。此部分提供指向 Etherscan 的链接或各种 Aave V3 操作的原始交易哈希值（例如，“供应 rETH 交易”、“借入 DAI 交易”、“提取 ETH 交易”）。

您可以复制这些交易哈希并将其粘贴到 Tenderly 中（如前所述）。分析这些真实交易将提供宝贵的背景信息和实际示例，帮助您理解协议的机制并顺利完成课程练习。

主要资源和工具摘要
总结一下，这些是你在整个 DeFi Aave V3 课程中将使用的基本资源和工具：

主要课程库：（ Cyfrin/defi-aave-v3包含练习、笔记、解决方案）。

官方 Aave V3 合约：（ github.com/aave-dao/aave-v3-origin具体标签v3.3.0）。

智能合约开发框架： Foundry（github.com/foundry-rs/foundry）。

以太坊主网 RPC 提供商： Alchemy.com（或任何其他提供商）。

交易调试工具： Tenderly.co。

Python 模拟环境（可选）： Jupyter Notebook。

设置好环境并了解这些资源后，您现在就可以使用 Aave V3 进入 DeFi 世界了！
