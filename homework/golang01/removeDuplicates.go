package golang01

import (
	"fmt"
)

func RemoveDuplicates(nums []int) int {

	var arr []int
	var numMap map[int]int = make(map[int]int)
	for _, v := range nums {
		if _, ok := numMap[v]; ok {
			continue
		} else {
			arr = append(arr, v)
			numMap[v] = 1
		}
	}

	fmt.Println("-------------------------arr=", arr)
	nums = arr
	return len(nums)

}
