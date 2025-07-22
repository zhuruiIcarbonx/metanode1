package main

import (
	"fmt"

	homework1 "github.com/zhuruiIcarbonx/metanode1/homework/golang01"

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
	homework1.IsPalindrome(121)
}
