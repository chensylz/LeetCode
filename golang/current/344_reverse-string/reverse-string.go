package _344_reverse_string

func reverseString(s []byte) {
	n := len(s)
	if n == 0 || n == 1 {
		return
	}
	for left, right := 0, n-1; left < right; {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
