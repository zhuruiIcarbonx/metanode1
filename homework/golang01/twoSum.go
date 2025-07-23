package golang01

import (
	"fmt"
)

/**
*** letcode执行时需删除所有fmt.Println。避免输出过长
**/
func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			fmt.Println("i=", i, "j=", j)
			if target == nums[i]+nums[j] {
				fmt.Println("return------i=", i, "j=", j)
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}

}
