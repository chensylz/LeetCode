package _6_permutations

var result [][]int

func permute(nums []int) [][]int {
	result = make([][]int, 0)
	if len(nums) == 0 {
		return result
	}
	p := make([]int, 0)
	used := make([]bool, len(nums))
	helper(nums, 0, p, used)
	return result
}

func helper(nums []int, index int, p []int, used []bool) {
	if index == len(nums) {
		tmp := make([]int, len(p))
		copy(tmp, p)
		result = append(result, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		used[i] = true
		p = append(p, nums[i])
		helper(nums, index+1, p, used)
		used[i] = false
		p = p[:len(p)-1]
	}
}
