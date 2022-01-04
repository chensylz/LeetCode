package _704_binary_search

import "testing"

type Result struct {
	Nums   []int
	Need   int
	Target int
}

func TestSearch(t *testing.T) {
	testMap := []Result{
		{
			[]int{-1, 0, 3, 5, 9, 12},
			4,
			9,
		},
		{
			[]int{-1, 0, 3, 5, 9, 12},
			-1,
			2,
		},
		{
			[]int{5},
			0,
			5,
		},
	}
	for _, value := range testMap {
		result := search(value.Nums, value.Target)
		if result != value.Need {
			t.Errorf("need %d, actual %d", value.Need, result)
		}
	}
}
