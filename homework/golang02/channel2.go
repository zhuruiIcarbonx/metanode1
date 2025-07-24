package golang02

import (
	"fmt"
	"time"
)

/**
* 2.题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
    考察点 ：通道的缓冲机制。
*/

func Chanel2() {
	ch := make(chan int, 100)

	// 生产者协程
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Printf(" I am sending : %d\n", i)
			ch <- i
		}
		close(ch) // 关闭通道
	}()

	// 消费者协程
	go func() {
		for num := range ch {
			fmt.Printf(" I received: %d\n", num)
		}
	}()

	// 等待一段时间以确保所有数据都被处理
	time.Sleep(2 * time.Second)
}
