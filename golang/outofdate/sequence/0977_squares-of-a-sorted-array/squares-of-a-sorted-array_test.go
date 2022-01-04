package _977_squares_of_a_sorted_array

import (
	"fmt"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	testValues := []struct {
		Nums []int
		Res  []int
	}{
		{[]int{-4, -1, 0, 3, 10},
			[]int{0, 1, 9, 16, 100},
		},
		{
			[]int{16, 1, 0, 9, 10},
			[]int{0, 1, 9, 16, 100},
		},
	}
	for _, value := range testValues {
		res := sortedSquares(value.Nums)
		fmt.Println(res)
		fmt.Println(value.Res)
	}
}
