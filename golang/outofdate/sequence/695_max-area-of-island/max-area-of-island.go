package _95_max_area_of_island


func maxAreaOfIsland(grid [][]int) int {
	var startPoints [][]int
	x, y := len(grid), len(grid[0])

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if grid[i][j] == 1 {
				startPoints = append(startPoints, []int{i, j})
			}
		}
	}

	if len(startPoints) == 0 {
		return 0
	}

	maxArea := 0

	for _, point := range startPoints {
		area := calculateIsland(grid, point[0], point[1])
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

var (
	dx = []int{0, 1, 0, -1}
	dy = []int{-1, 0, 1, 0}
)

func calculateIsland(grid [][]int, tx int, ty int) int {
	area := 1
	var queue [][]int
	queue = append(queue, []int{tx, ty})
	grid[tx][ty] = 0
	x, y := len(grid), len(grid[0])
	for i := 0; i < len(queue); i++ {
		current := queue[i]
		for j := 0; j < 4; j++ {
			cx, cy := current[0] + dx[j], current[1] + dy[j]
			if cx >= 0 && cy >= 0 && cx < x && cy < y && grid[cx][cy] == 1 {
				area++
				queue = append(queue, []int{cx, cy})
				grid[cx][cy] = 0
			}
		}
	}
	return area
}