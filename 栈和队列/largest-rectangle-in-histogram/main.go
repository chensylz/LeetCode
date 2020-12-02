package main

/**
84. 柱状图中最大的矩形
https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
 */


func largestRectangleArea(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack) - 1]] >= heights[i] {
			// 当它被弹出时，代表着它的右边没有柱子比它高了
			right[stack[len(stack) - 1]] = i
			stack = stack[:len(stack) - 1]
		}

		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack) - 1]
		}
		stack = append(stack, i)
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans = Max(ans, (right[i] - left[i] - 1) * heights[i])
	}
	return ans
}


func largestRectangleArea2(heights []int) int {
	n := len(heights)
	left := make([]int, n)
	right := make([]int, n)
	stack := make([]int, 0)

	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack) - 1]] >= heights[i] {
			stack = stack[:len(stack) - 1]
		}

		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack) - 1]
		}
		stack = append(stack, i)
	}

	stack = make([]int, 0)

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[stack[len(stack) - 1]] >= heights[i] {
			stack = stack[:len(stack) - 1]
		}

		if len(stack) == 0 {
			right[i] = n
		} else {
			right[i] = stack[len(stack) - 1]
		}
		stack = append(stack, i)
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans = Max(ans, (right[i] - left[i] - 1) * heights[i])
	}
	return ans
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	largestRectangleArea([]int{2,1,5,6,2,3})
}