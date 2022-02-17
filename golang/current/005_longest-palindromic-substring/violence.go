package _05_longest_palindromic_substring

func longestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	maxLen := 1
	begin := 0

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if j-i+1 > maxLen && checkIsPalindrome(s, i, j) {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}

func checkIsPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
