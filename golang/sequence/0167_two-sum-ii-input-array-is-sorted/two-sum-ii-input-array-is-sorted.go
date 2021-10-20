package _167_two_sum_ii_input_array_is_sorted


func twoSum(numbers []int, target int) []int {
	res := make(map[int]int)
	for index, value := range numbers {
		if v, ok := res[value]; ok {
			return []int{v, index + 1}
		}
		res[target - value] = index + 1
	}
	return nil
}
