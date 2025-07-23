package golang01

import "strings"

/****
最长公共前缀

编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""
**/

func LongestCommonPrefix(strs []string) string {

	prifix := strs[0]
	for i := len(prifix) - 1; i >= 0; i-- {
		for j := 0; j < len(strs); j++ {
			str := strs[j]
			if len(str) >= len(prifix) && strings.HasPrefix(str, prifix) {
				if j == len(strs)-1 {
					return prifix
				}
				continue
			} else {
				prifix = prifix[:i-1]
				break
			}

		}
	}
	return ""
}
