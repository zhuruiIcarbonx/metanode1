package main

import (
	"fmt"

	gorm "github.com/zhuruiIcarbonx/metanode1/homework/gorm"
	lesson1 "github.com/zhuruiIcarbonx/metanode1/lesson1"
)

func main() {
	a := lesson1.Add(1, 2)
	fmt.Println("Hello, World!", a)

	// lesson14.DoSth()
	// lesson14.UseMap()
	// lesson14.MapLock()
	// lesson15.RangeStr()
	// lesson15.RangeCN()
	// lesson15.RangeSliceAndArray()
	// lesson15.RangeChan()

	// lesson16.DoInterface()
	// homework1.LongestCommonPrefix([]string{"flower", "flow", "flight"})

	// arr := homework1.PlusOne([]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6})

	// arr := homework1.Merge([][]int{{2, 3}, {2, 2}, {3, 3}, {1, 3}, {5, 7}, {2, 2}, {4, 6}})

	// fmt.Println("-------------------------arr=", arr)

	// homework2.Chanel1()
	// homework2.Chanel2()
	// homework2.Goroutine1()
	// homework2.Goroutine2()
	// homework2.Pointer1()
	// homework2.Pointer2()
	// homework2.Object1()
	// homework2.Object2()
	// homework2.Lock1()
	// homework2.Lock2()
	// gorm.Crud()
	gorm.Transfer()

}
