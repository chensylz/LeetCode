package Tools

import "math/rand"

func GenerateRandArray(size int) []int {
	result := make([]int, 0)
	for i := 0; i < size; i++ {
		result = append(result, rand.Intn(100))
	}
	return result
}