package evaluate_reverse_polish_notation

import "strconv"

/**
150. 逆波兰表达式求值
https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/
 */

func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return -1
	}
	stack := make([]int, 0)

	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+", "-", "*", "/":
			if len(stack) < 2 {
				return -1
			}

			//逆序取出
			b := stack[len(stack) - 1]
			a := stack[len(stack) - 2]
			stack = stack[:len(stack) - 2]
			var result int

			switch tokens[i] {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				result = a / b
			}
			stack = append(stack, result)
		default:
			val, _ := strconv.Atoi(tokens[i])
			stack = append(stack, val)
		}
	}
	return stack[0]
}
