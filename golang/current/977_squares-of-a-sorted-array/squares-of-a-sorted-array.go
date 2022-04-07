package main

import "fmt"

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}

func sortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, 0, n)
	mid := 0
	for i, v := range nums {
		if v >= 0 {
			mid = i
			break
		}
	}
	start, end := mid-1, mid
	if mid == 0 && nums[mid] < 0 {
		start, end = n-1, n
	} else if mid == 0 && nums[mid] >= 0 {
		start, end = -1, 0
	}

	for start >= 0 || end < n {
		if start < 0 {
			result = append(result, nums[end]*nums[end])
			end += 1
		} else if end >= n {
			result = append(result, nums[start]*nums[start])
			start -= 1
		} else if nums[start]*nums[start] < nums[end]*nums[end] {
			result = append(result, nums[start]*nums[start])
			start -= 1
		} else {
			result = append(result, nums[end]*nums[end])
			end += 1
		}
	}
	return result
}
