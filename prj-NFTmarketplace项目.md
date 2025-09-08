## 下周学习内容

    1、开始学习NFT marketplace项目
    2、先看视频和文档，理解整体架构和代码
    3、然后阅读合约以及后端代码
    
    参考链接：https://github.com/MetaNodeAcademy/ProjectBreakdown-NFTMarket/blob/main/NFTMartket.md
    
    重点是 EasySwapSync模块以及NFT Market项目面试指南，带着里面的问题去阅读代码
    
    重中之重是NFT Market项目面试指南
    
    面试必问：链上链下数据同步(章节)



## js用法技巧：
### 1、callStatic
    orderKeys = await esDex.callStatic.makeOrders(orders)
    expect(orderKeys[0]).to.not.equal(Byte32Zero)

### 2、attach
    let testERC721Address = "0xF2e0BA02a187F19F5A390E4f990c684d81A833A0";
    let testERC721 = await (await ethers.getContractFactory("Troll")).attach(testERC721Address)
    tx = await testERC721.mint(deployer.address, 50);
    await tx.wait()
    console.log("mint tx:", tx.hash)

## 项目工具链接：

    http://evm.codes/
	
	slither作用：1、检测合约漏洞  
	https://github.com/crytic/slither


## 项目命令功能
### 1. Get Contract Size
    npx hardhat size-contracts


### 2. see storage layout of contract
    查看每个变量占用哪个slot
    slither-read-storage ./contracts/EasySwapOrderBook.sol --contract-name EasySwapOrderBook --solc-remaps @=node_modules/@ --json storage_layout.json

