package leetcode_go

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	ss := strings.Fields(s)
	return len(ss[len(ss)-1])
}
