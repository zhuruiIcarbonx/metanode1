package golang02

import (
	"fmt"
	"time"
)

/***

1.题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。

*/

func Goroutine1() {

	ch := make(chan int)

	go func() {
		for i := 1; i < 10; i += 2 {
			ch <- i
			fmt.Println("ch1 is print:", i)

		}
	}()

	go func() {
		for i := 2; i < 10; i += 2 {
			r, _ := <-ch
			fmt.Println("ch2 is print:", i, "----r=", r)

		}
	}()

	for i := 0; i < 5; i++ {

		fmt.Println("main is print:", i)
		time.Sleep(time.Second) // Sleep to allow goroutines to run
	}
	close(ch)
}
