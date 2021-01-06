package main

import (
	"sort"
)

type Intvs [][]int

func (i Intvs) Len() int {
	return len(i)
}

func (i Intvs) Swap(j, k int) {
	i[j][0], i[j][1] = i[k][0], i[k][1]
}

func (i Intvs) Less(j, k int) bool {
	return i[j][1] < i[k][1]
}


func intervalSchedule(intvs [][]int) int {
	if len(intvs) == 0 {
		return 0
	}
	var intvsSelf Intvs
	intvsSelf = intvs
	sort.Sort(intvsSelf)

	// 至少有一个区间不相交
	count := 1
	xEnd := intvsSelf[0][1]
	for _, interval := range intvsSelf {
		start := interval[0]
		if start >= xEnd {
			count++
			xEnd = interval[1]
		}
	}
	return count
}

func main() {
	intervalSchedule([][]int{{1, 3}, {2, 4}, {3, 6}})
}

