package _33_flood_fill

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	queue := [][]int{}
	x, y := len(image), len(image[0])
	if sr >= x || sc >= y {
		return image
	}
	currentColor := image[sr][sc]
	if currentColor == newColor {
		return image
	}
	queue = append(queue, []int{sr, sc})
	image[sr][sc] = newColor

	var (
		dx = []int{0, 1, 0, -1} // 上右下左
		dy = []int{-1, 0, 1, 0} // 上右下左
	)


	for i := 0; i < len(queue); i++ {
		current := queue[i]
		for j := 0; j < 4; j++ {
			ux, uy := current[0] + dx[j], current[1] + dy[j]
			if ux >= 0 && uy >= 0 && ux < x && uy < y && image[ux][uy] == currentColor {
				queue = append(queue, []int{ux, uy})
				image[ux][uy] = newColor
			}
		}

	}

	return image
}