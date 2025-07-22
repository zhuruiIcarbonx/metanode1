package lesson15

import (
	"fmt"
	"reflect"
	"time"
)

// 这里需要强调一点，在 Go 中，所有字符串都是按照 Unicode 编码的。
func RangeStr() {
	str1 := "abc123"
	for index := range str1 {
		fmt.Printf("str1 -- index:%d, value:%d\n", index, str1[index])
	}

	str2 := "测试中文"
	for index := range str2 {
		fmt.Printf("str2 -- index:%d, value:%d\n", index, str2[index])
	}
	fmt.Printf("len(str2) = %d\n", len(str2))

	runesFromStr2 := []rune(str2)
	bytesFromStr2 := []byte(str2)
	fmt.Printf("len(runesFromStr2) = %d\n", len(runesFromStr2))
	fmt.Printf("len(bytesFromStr2) = %d\n", len(bytesFromStr2))
}

func RangeCN() {
	str1 := "a1中文"
	for index, value := range str1 {
		fmt.Printf("str1 -- index:%d, index value:%d\n", index, str1[index])
		fmt.Printf("str1 -- index:%d, range value:%d\n", index, value)
	}
}

func RangeSliceAndArray() {
	array := [...][3]int{{1, 2, 3}, {4, 5, 6}}
	slice := [][]int{{1, 2}, {3}}
	// 只拿到行的索引
	for index := range array {
		// array[index]类型是一维数组
		fmt.Println(reflect.TypeOf(array[index]))
		fmt.Printf("array -- index=%d, value=%v\n", index, array[index])
	}

	for index := range slice {
		// slice[index]类型是一维数组
		fmt.Println(reflect.TypeOf(slice[index]))
		fmt.Printf("slice -- index=%d, value=%v\n", index, slice[index])
	}

	// 拿到行索引和该行的数据
	fmt.Println("print array element")
	for row_index, row_value := range array {
		fmt.Println(row_index, reflect.TypeOf(row_value), row_value)
	}

	fmt.Println("print array slice")
	for row_index, row_value := range slice {
		fmt.Println(row_index, reflect.TypeOf(row_value), row_value)
	}

	// 双重遍历，拿到每个元素的值
	for row_index, row_value := range array {
		for col_index, col_value := range row_value {
			fmt.Printf("array[%d][%d]=%d ", row_index, col_index, col_value)
		}
		fmt.Println()
	}
	for row_index, row_value := range slice {
		for col_index, col_value := range row_value {
			fmt.Printf("slice[%d][%d]=%d ", row_index, col_index, col_value)
		}
		fmt.Println()
	}
}

func addData(ch chan int) {
	size := cap(ch)
	for i := 0; i < size; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
	close(ch)
}

func RangeChan() {
	ch := make(chan int, 10)

	go addData(ch)

	for i := range ch {
		fmt.Println(i)
	}
}
