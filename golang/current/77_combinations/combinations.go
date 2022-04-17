package _7_combinations

var res [][]int

func combine(n int, k int) [][]int {
	res = [][]int{}
	if n <= 0 || k <= 0 || n < k {
		return res
	}
	helper(n, k, 1, []int{})
	return res
}

func helper(n int, k int, start int, tracker []int) {
	if len(tracker) == k {
		temp := make([]int, k)
		copy(temp, tracker)
		res = append(res, temp)
		return
	}
	// 如果剩余的数量已不足满足K长度，那么没必要继续找了
	if len(tracker)+n-start+1 < k {
		return
	}
	for i := start; i <= n; i++ {
		tracker = append(tracker, i)
		helper(n, k, i+1, tracker)
		tracker = tracker[:len(tracker)-1]
	}
}
