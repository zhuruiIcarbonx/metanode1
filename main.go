package main

import (
	"fmt"

	lesson1 "github.com/zhuruiIcarbonx/metanode1/lesson1"
	lesson14 "github.com/zhuruiIcarbonx/metanode1/lesson14"
)

func main() {
	a := lesson1.Add(1, 2)
	fmt.Println("Hello, World!", a)

	// lesson14.DoSth()
	// lesson14.UseMap()
	lesson14.MapLock()
}
