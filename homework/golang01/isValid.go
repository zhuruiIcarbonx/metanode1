package golang01

/**

有效的括号

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
*/

func IsValid(s string) bool {

	stack := make([]rune, 0)
	for _, v := range s {
		switch v {
		case '(', '[', '{':
			stack = append(stack, v)
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[0 : len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[0 : len(stack)-1]
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[0 : len(stack)-1]

		}
	}

	if len(stack) > 0 {
		return false
	}

	return true

}
