package triangle

/**
120. 三角形最小路径和
https://leetcode-cn.com/problems/triangle/
 */
var memory [][]int

func minimumTotal2(triangle [][]int) int {
	memory = make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		memory[i] = make([]int, len(triangle[len(triangle) - 1]))
	}
	return recursion(triangle, 0, 0)
}



func recursion(triangle [][]int, i, j int) int {
	if i == len(triangle) {
		return 0
	}

	if memory[i][j] != 0 {
		return memory[i][j]
	}
	memory[i][j] = min(recursion(triangle, i + 1, j), recursion(triangle, i + 1, j + 1)) + triangle[i][j]
	return memory[i][j]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
