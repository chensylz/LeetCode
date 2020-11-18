package main

import "fmt"

/**
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
https://leetcode-cn.com/problems/subsets/
 */

func subsets(nums []int) [][]int {
	result := make([][]int, 0)

	list := make([]int, 0)

	backtrack(nums, 0, list, &result)

	return result
}


func backtrack(nums []int, pos int, list []int, result *[][]int) {

	ans := make([]int, len(list))
	copy(ans, list)
	*result = append(*result, ans)

	for i := pos; i < len(nums); i++ {
		list = append(list, nums[i])
		backtrack(nums, i + 1, list, result)
		list = list[0 : len(list) - 1]
	}
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3, 4, 5}))
}
