package sortedfunction

// Shell 希尔排序.
func Shell(nums []int) {
	for gap := len(nums) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(nums); i++ {
			for j := i - gap; j >= 0 && nums[j] > nums[j+gap]; j -= gap {
				nums[j], nums[j+gap] = nums[j+gap], nums[j]
			}
		}
	}
}
