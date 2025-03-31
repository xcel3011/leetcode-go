package leetcode_go

import (
	"strings"
)

func addSpaces(s string, spaces []int) string {
	b := strings.Builder{}
	idx := 0
	for i := 0; i < len(s); i++ {
		if idx < len(spaces) && i == spaces[idx] {
			b.WriteByte(' ')
			idx++
		}
		b.WriteByte(s[i])
	}
	return b.String()
}
