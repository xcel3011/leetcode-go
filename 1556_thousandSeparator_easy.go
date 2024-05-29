package leetcode_go

import (
	"strconv"
)

func thousandSeparator(n int) string {
	return solve(strconv.Itoa(n))
}

func solve(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return solve(s[:n-3]) + "." + s[n-3:]
}
