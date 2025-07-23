package golang02

import (
	"fmt"
	"time"
)

/**
 * 1.题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来
 */

func Chanel1() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for num := range ch {
		fmt.Println(num)
	}
}
