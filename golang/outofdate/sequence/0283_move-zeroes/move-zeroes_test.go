package _283_move_zeroes

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	testValues := []struct {
		Nums []int
		Res  []int
	}{
		{[]int{0,1,0,3,12},
			[]int{1,3,12,0,0},
		},
		{
			[]int{1, 0, 1},
			[]int{1, 1, 0},
		},
	}
	for _, value := range testValues {
		moveZeroes(value.Nums)
		fmt.Println("结果:", value.Nums)
		fmt.Println("预期:", value.Res)
	}
}