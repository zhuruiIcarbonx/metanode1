## 实践任务：

    1、部署一个带mint和burn功能的erc20合约，铸造销毁几个token，转移几个token，来构造事件
    2、使用go语言写一个后端服务来追踪合约事件，重建用户的余额
    3、以太坊延迟六个区块，确保区块链不会回滚
    4、加上积分计算功能，起一个定时任务，每小时根据用户的余额来计算用户的积分，暂定积分是余额*0.05
    5、要记录用户的所有余额变化，根据这个来计算积分，这样更准确一些
    6、需要维护一下用户的总余额表以及总积分表，还有一个用户的余额变动记录表
    7、需要支持多链逻辑，比如支持sepolia， base sepolia




```c
    举个例子：
    用户在15：00的时候0个token，15：10分有100个token，15：30有200个token
    计算积分的时候，需要考虑用户的余额变化
    
    比如此时是16：00启动定时任务了来计算积分，应该是100*0.05*20/60+200*0.05*30/60
    
```
考虑一个场景，如果程序错误了，或者rpc有问题，导致好几天没有计算积分。此时应该如何正确回溯？

    https://sepolia.basescan.org/
    https://superbridge.app/base-sepolia
    https://github.com/DCreek03/web3.0-learning/tree/main/go_work/task4

## markdown online
    https://zlsam.com/markdown/
	
