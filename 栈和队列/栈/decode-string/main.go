package decode_string

import "strconv"

/**
394. 字符串解码
https://leetcode-cn.com/problems/decode-string/
 */

func decodeString(s string) string {
	if s == "" {
		return s
	}
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ']':
			// 保存括号内的
			temp := make([]byte, 0)

			for len(stack) > 0 && stack[len(stack) - 1] != '[' {
				temp = append(temp, stack[len(stack) - 1])
				stack = stack[:len(stack) - 1]
			}

			// 取出 '['
			stack = stack[:len(stack) - 1]

			// 找到计数的长度
			idx := 1
			for len(stack) >= idx && stack[len(stack) - idx] >= '0' && stack[len(stack) - idx] <= '9' {
				idx++
			}

			num := stack[len(stack) - idx + 1:]
			stack = stack[:len(stack) - idx + 1]

			count, _ := strconv.Atoi(string(num))

			// 重复放入栈中
			for j := 0; j < count; j++ {
				// 注意，逆序取的逆序放
				for z := len(temp) - 1; z >= 0; z-- {
					stack = append(stack, temp[z])
				}
			}
		default:
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}