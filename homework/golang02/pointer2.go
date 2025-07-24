package golang02

import "fmt"

/***

2.题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
考察点 ：指针运算、切片操作。
*/

func Pointer2() {
	arr := []int{1, 2, 3, 4, 5}
	p := &arr
	multi(p)
	fmt.Printf("Modified slice: %v\n", *p)

}

func multi(arr *[]int) {

	for i := range *arr {
		(*arr)[i] *= 2
	}

}
