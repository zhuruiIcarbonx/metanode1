package golang02

import (
	"fmt"
	"time"
)

/***

2.题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。

*/

func Goroutine2() {

	tasks := []func(){

		func() { time.Sleep(2 * time.Second); fmt.Println("task 1 completed") },
		func() { time.Sleep(1 * time.Second); fmt.Println("task 2 completed") },
		func() { time.Sleep(1 * time.Second); fmt.Println("task 3 completed") },
	}

	for _, task := range tasks {
		go func(t func()) {
			start := time.Now()
			t()
			duration := time.Since(start)
			fmt.Printf("Task executed in %v\n", duration)
		}(task)
	}

	time.Sleep(5 * time.Second)

}
