package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	left, right, l := 0, 0, 0
	for i := 1; i < n; i++ {
		if strings.Contains(s[left:i], string(s[i])) {
			left = strings.Index(s[left:i], string(s[i])) + 1 + left
		}
		right = i
		if l < right-left+1 {
			l = right - left + 1
		}
	}
	return l
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
