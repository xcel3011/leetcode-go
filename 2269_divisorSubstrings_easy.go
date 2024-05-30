package leetcode_go

import (
	"strconv"
)

func divisorSubstrings(num int, k int) int {
	f := func(n int) bool {
		return n != 0 && num%n == 0
	}

	s := strconv.Itoa(num)
	slow, fast := 0, k
	count := 0
	for fast <= len(s) {
		n, _ := strconv.Atoi(s[slow:fast])
		if f(n) {
			count++
		}
		slow++
		fast++
	}
	return count
}
