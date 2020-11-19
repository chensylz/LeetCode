package SoredFunction


func ShellSort(nums []int) []int {
	n := len(nums)

	step := 1
	for step < n / 3 {
		step = 3 * step + 1
	}

	for step > 0 {
		for index := step; index < n; index++ {
			for j := index; j >= step && nums[j] < nums[j - step]; j -= step {
				nums[j], nums[j - step] =  nums[j - step],  nums[j]
			}
		}
		step = step / 3
	}

	return nums
}