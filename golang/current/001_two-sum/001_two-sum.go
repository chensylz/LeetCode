package _01_two_sum

func twoSum(nums []int, target int) []int {
	targetMap := make(map[int]int)
	for i, num := range nums {
		if j, ok := targetMap[target-num]; ok {
			return []int{i, j}
		}
		targetMap[num] = i
	}
	return nil
}
