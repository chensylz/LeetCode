package 快速排序


func QuickSort(nums []int) []int {
	//  思路：把一个数组分为左右两段，左段小于右段
	quickSort(nums, 0, len(nums) - 1)
	return nums
}

func quickSort(nums []int, start, end int) {
	if start < end {
		flag := partition(nums, start, end)
		quickSort(nums, 0, flag-1)
		quickSort(nums, flag + 1, end)
	}
}

func partition(nums []int, start, end int) int {
	// 选最后一个
	flag := nums[end]
	i := start
	for j := start; j < end; j++ {
		if nums[j] < flag {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[end] = nums[end], nums[i]
	return i
}