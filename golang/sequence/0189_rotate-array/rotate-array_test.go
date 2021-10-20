package _189_rotate_array

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	testValues := []struct {
		Nums []int
		Res  []int
		K    int
	}{
		{[]int{1,2,3,4,5,6,7},
			[]int{5,6,7,1,2,3,4},
			3,
		},
		{
			[]int{-1,-100,3,99},
			[]int{3,99,-1,-100},
			2,
		},
	}
	for _, value := range testValues {
		rotate(value.Nums, value.K)
		fmt.Println("结果:", value.Nums)
		fmt.Println("预期:", value.Res)
	}
}