package _003_longest_substring_without_repeating_characters


func lengthOfLongestSubstring(s string) int {
	winStr := make([]byte, 0)
	maxLen := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if index := findIndex(winStr, c); index != -1 {
			if n := len(winStr); n > maxLen {
				maxLen = n
			}
			winStr = winStr[index + 1:]
		}
		winStr = append(winStr, c)
	}
	if n := len(winStr); n > maxLen {
		maxLen = n
	}
	return maxLen
}

func findIndex(b []byte, target byte) int {
	for index, v := range b {
		if v == target {
			return index
		}
	}
	return -1
}