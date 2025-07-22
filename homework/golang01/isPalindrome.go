package golang01

import (
	"fmt"
	"strconv"
)

func IsPalindrome(x int) bool {

	str := strconv.Itoa(x)
	length := len(str)

	for index := range str {
		if index >= length/2+1 {
			break
		}
		fmt.Println("当前index:", index)

		if str[index] == str[length-index-1] {

		} else {
			fmt.Println(x, "不是回文数")
			return false
		}
	}
	fmt.Println(x, "是回文数")
	return true

}
