package main

import (
	"fmt"
	"sort"
)

/**
435. 无重叠区间
https://leetcode-cn.com/problems/non-overlapping-intervals/
 */


// 动态规划
func eraseOverlapIntervalsV1(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			// 如果不在区间内
			if intervals[j][1] <= intervals[i][0] {
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
	}
	return n - max(dp...)
}

func max(x ...int) int {
	res := x[0]
	for _, v := range x {
		if v > res {
			res = v
		}
	}
	return res
}

// 贪心算法
func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	count := 1
	end := intervals[0][1]
	for _, inter := range intervals[1:]{
		if inter[0] >= end {
			count++
			end = inter[1]
		}
	}
	return n - count
}

func main() {
	fmt.Print(eraseOverlapIntervals([][]int{{0,1},{3,4},{1,2}}))
}