package _035_search_insert_position

import "testing"

type Result struct {
	Nums   []int
	Need   int
	Target int
}

func TestSearch(t *testing.T) {
	testMap := []Result{
		{
			[]int{1,3,5,6},
			2,
			5,
		},
		{
			[]int{1,3,5,6},
			1,
			2,
		},
		{
			[]int{1,3,5,6},
			4,
			7,
		},
	}
	for _, value := range testMap {
		result := searchInsert(value.Nums, value.Target)
		if result != value.Need {
			t.Errorf("need %d, actual %d", value.Need, result)
		}
	}
}
