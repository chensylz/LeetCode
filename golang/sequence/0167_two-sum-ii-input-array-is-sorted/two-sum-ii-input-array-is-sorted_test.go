package _167_two_sum_ii_input_array_is_sorted

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testValues := []struct {
		Nums []int
		Res  []int
		K    int
	}{
		{[]int{2,7,11,15},
			[]int{1,2},
			9,
		},
	}
	for _, value := range testValues {
		res := twoSum(value.Nums, value.K)
		fmt.Println("结果:", res)
		fmt.Println("预期:", value.Res)
	}
}