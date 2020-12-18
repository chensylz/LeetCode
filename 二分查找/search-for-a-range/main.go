package search_for_a_range


/**
61. 搜索区间
https://www.lintcode.com/problem/search-for-a-range/description
 */

func searchRange (A []int, target int) []int {
	if len(A) == 0 {
		return []int{-1, -1}
	}
	result := make([]int, 2)
	start := 0
	end := len(A) - 1
	for start + 1 < end {
		mid := start + (end - start) / 2
		if A[mid] > target {
			end = mid
		} else if A[mid] < target {
			start = mid
		} else {
			// 如果相等，应该继续向左找，就能找到第一个目标值的位置
			end = mid
		}
	}

	if A[start] == target {
		result[0] = start
	} else if A[end] == target {
		result[0] = end
	} else {
		result[0], result[1] = -1, -1
		return result
	}

	start = 0
	end = len(A) - 1
	for start + 1 < end {
		mid := start + (end - start) / 2
		if A[mid] > target {
			end = mid
		} else if A[mid] < target {
			start = mid
		} else {
			// 如果相等，应该继续向右找，就能找到最后一个目标值的位置
			start = mid
		}
	}

	if A[end] == target {
		result[1] = end
	} else if A[start] == target {
		result[1] = start
	} else {
		result[1], result[1] = -1, -1
		return result
	}
	return result
}