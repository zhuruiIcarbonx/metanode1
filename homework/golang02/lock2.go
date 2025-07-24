package golang02

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/**
*2.题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全。
*/

func Lock2() {
	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
			wg.Done() //这一步要在协程里面执行
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}
