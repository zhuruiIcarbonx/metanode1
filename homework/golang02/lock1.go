package golang02

/**
*1.题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。
*/

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

func Lock1() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
			wg.Done() //这一步要在协程里面执行
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}
