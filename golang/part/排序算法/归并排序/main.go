package 归并排序

func MerageSort(nums []int) []int {
	return mergeSort(nums)
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	// 分治，一次分成两边
	mid := len(nums) / 2
	left := merageSort(nums[:mid])
	right := merageSort(nums[mid:])

	// 合并
	result := merge(left, right)
	return result
}

func merge(left, right []int) (result []int) {
	l := 0
	r := 0
	for l < len(left) && r < len(right) {
		// 判断谁更小
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}