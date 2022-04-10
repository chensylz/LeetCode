package _83_move_zeroes

func moveZeroes(nums []int) {
	n := len(nums)
	if n == 0 || n == 1 {
		return
	}
	for i, j := 0, 1; i < n && j < n; {
		if nums[i] == 0 && nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
		}
		if nums[i] != 0 {
			i++
		}
		j++
	}
}
