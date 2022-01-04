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

func towSum2(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		l, r := i + 1, len(numbers) - 1
		for l <= r {
			mid := (r - l) / 2 + l
			if numbers[mid] == target - numbers[i] {
				return []int{i, mid}
			} else if numbers[mid] > target - numbers[i] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return nil
}
