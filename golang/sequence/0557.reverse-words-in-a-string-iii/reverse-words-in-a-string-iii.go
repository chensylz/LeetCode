package _57_reverse_words_in_a_string_iii

import "strings"

func reverseWords(s string) string {
	sArr := strings.Split(s, " ")
	n := len(sArr)
	sRes := make([]string, n)
	for index, str := range sArr {
		sRes[n - index - 1] = reverse(str)
	}
	return strings.Join(sRes, " ")
}

func reverse(s string) string {
	sBytes := []byte(s)
	n := len(s)
	for i := 0; i < n / 2; i++ {
		sBytes[i], sBytes[n - i - 1] = sBytes[n - i - 1], sBytes[i]
	}
	return string(sBytes)
}