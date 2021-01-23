package minimum_number_of_arrows_to_burst_balloons

import "sort"

/**
452. 用最少数量的箭引爆气球
https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/
 */

func findMinArrowShots(points [][]int) int {
	n := len(points)
	if n == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	count := 1
	end := points[0][1]
	for _, point := range points{
		if point[0] > end {
			count ++
			end = point[1]
		}
	}
	return count
}