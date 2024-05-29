package leetcode_go

import "strings"

func isPalindrome(s string) bool {
	var ss string
	s = strings.ToLower(s)
	for i := range s {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= '0' && s[i] <= '9') {
			ss += string(s[i])
		}
	}

	slow, fast := 0, len(ss)-1
	for fast >= slow {
		if ss[slow] != ss[fast] {
			return false
		}
		slow++
		fast--
	}

	return true
}
