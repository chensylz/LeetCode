package _57_reverse_words_in_a_string_iii

import "strings"

func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	result := make([]string, len(arr))
	for i, a := range arr {
		result[i] = reverseString([]byte(a))
	}
	return strings.Join(result, " ")
}

func reverseString(s []byte) string {
	for start, end := 0, len(s)-1; start < end; {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
	return string(s)
}
