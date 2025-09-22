# golang中case后面为什么会出现“<-”

    在Golang 的 select 语句的 case 后面出现"<-" 是因为该 case 属于 channel 操作，它表示对一个channel 进行发送或接收操作，表示接收（`<-ch）或发送（ch<-`value）数据。﻿
    
	接收操作:case <-ch: 表示等待从channel ch 中接收数据。当 ch 中有数据时，该case 就会被执行。
    发送操作:case ch <- value: 表示将 value 发送至channel ch。当 ch 有可读的接收者时，该case 就会被执行。
    
	为什么会这样设计？
    Go 语言的 select 语句是用于处理多个channel 的并发操作的。它允许程序在多个channel 操作中选择一个就绪的进行执行，避免阻塞在某个单一的channel 操作上。﻿
    简洁的语法:"<-" 符号简洁地表达了channel 的读写方向，使得代码更加易读和直观。﻿
    并发处理:结合 select 语句和channel 的"<-" 操作，Go 语言能够有效地处理并发任务，例如在网络编程、分布式系统和微服务中。﻿
	
    举例说明
    下面是一个使用 select 和 case 处理channel 的简单示例:﻿
	
	package main
    import "fmt"
    import "time"
    
    func main() {
    	c1 := make(chan string)
    	c2 := make(chan string)
    
    	go func() {
    		time.Sleep(time.Second * 1)
    		c1 <- "结果 1"
    	}()
    	go func() {
    		time.Sleep(time.Second * 2)
    		c2 <- "结果 2"
    	}()
    
    	for i := 0; i < 2; i++ {
    		select {
    		case msg1 := <-c1:
    			fmt.Println("接收到:", msg1)
    		case msg2 := <-c2:
    			fmt.Println("接收到:", msg2)
    		}
    	}
    }
	
	

    在这个例子中，<-c1 和 <-c2 都表示从相应的channel 中接收数据。﻿

    总结
    "<-" 在Golang 的 case 语句中表示一个channel 的发送或接收操作，它是Go 语言中实现并发和channel 通信的核心语法
	
	
	
