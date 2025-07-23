package golang01

import (
	"fmt"
	"strings"
)

/*
***
最长公共前缀

编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""
*
*/
func longestCommonPrefix(strs []string) string {

	if len(strs) == 1 {
		return strs[0]
	}

	prifix := strs[0]
	for i := len(prifix) - 1; i >= 0; i-- {
		fmt.Println("prifix=", prifix, "i=", i)
		for j := 1; j < len(strs); j++ {

			fmt.Println("i=", i, "-----j=", j)
			str := strs[j]
			if len(str) >= len(prifix) && strings.HasPrefix(str, prifix) {
				if j == len(strs)-1 {
					return prifix
				}
				continue
			} else {
				break
			}
		}
		if i == 1 {
			prifix = prifix[0:1]
		} else if i == 0 {
			return ""
		} else {
			prifix = prifix[:i]
		}

	}
	return ""
}
