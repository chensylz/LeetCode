package _67_permutation_in_string


func checkInclusion(s1 string, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	var win1, win2 [26]int
	for i := 0; i < n; i++ {
		win1[s1[i] - 'a']++
		win2[s2[i] - 'a']++
	}

	if win1 == win2 {
		return true
	}

	for i := n; i < m; i++ {
		win2[s2[i] - 'a'] ++
		win2[s2[i - n] - 'a'] --
		if win1 == win2 {
			return true
		}
	}
	return false
}

