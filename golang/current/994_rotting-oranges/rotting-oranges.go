package main

import "fmt"

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	q := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				q = append(q, []int{i, j})
			}
		}
	}
	directions := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	level := -1
	for len(q) > 0 {
		tmp := make([][]int, 0)
		level++
		for i := 0; i < len(q); i++ {
			for _, d := range directions {
				x := q[i][0] + d[0]
				y := q[i][1] + d[1]
				if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
					grid[x][y] = 2
					tmp = append(tmp, []int{x, y})
				}
			}
		}
		q = tmp
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	if level == -1 {
		return 0
	}
	return level
}

func main() {
	fmt.Println(orangesRotting([][]int{
		{2, 1, 1},
		{0, 1, 1},
		{1, 0, 1},
	}))
}
