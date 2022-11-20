package tools

import "math/rand"

func GenerateRandomInts(size int) []int {
	ints := make([]int, size)
	for i := 0; i < size; i++ {
		ints[i] = rand.Intn(100)
	}
	return ints
}

func Copy(src []int) []int {
	actual := make([]int, len(src))
	copy(actual, src)
	return actual
}
