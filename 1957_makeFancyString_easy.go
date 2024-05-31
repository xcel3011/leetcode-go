package leetcode_go

import (
	"strings"
)

func makeFancyString(s string) string {
	ans := strings.Builder{}
	n := len(s)
	for i := 0; i < n; {
		start := i
		i++
		for i < n && s[i] == s[i-1] {
			i++
		}

		count := min(2, i-start)
		for count > 0 {
			ans.Write([]byte{s[start]})
			count--
		}
	}
	return ans.String()
}
