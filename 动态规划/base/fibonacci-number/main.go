package fibonacci_number

/**
509. 斐波那契数
https://leetcode-cn.com/problems/fibonacci-number/
 */

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	prev, curr := 1, 1
	for i := 3; i <= n; i++ {
		sum := prev + curr
		prev = curr
		curr = sum
	}
	return curr
}
