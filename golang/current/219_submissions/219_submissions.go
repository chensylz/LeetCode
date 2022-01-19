package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	for i := 0; i < n; i++ {
		var (
			start int
			end   = n
		)
		if i >= k {
			start = i - k
		}
		if i+k <= n {
			end = i + k
		}
		for j := start; j < end; j++ {
			if nums[i] == nums[j] && i != j {
				return true
			}
		}
	}
	return false
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
}
