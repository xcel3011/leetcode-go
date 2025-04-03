package leetcode_go

import "strings"

func isPalindrome125(s string) bool {
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

func isPalindrome250326(s string) bool {
	b := strings.Builder{}
	s = strings.ToLower(s)
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= '0' && s[i] <= '9') {
			b.WriteByte(s[i])
		}
	}

	ss := b.String()
	l, r := 0, len(ss)-1
	for l <= r {
		if ss[l] != ss[r] {
			return false
		}
		l++
		r--
	}
	return true
}
