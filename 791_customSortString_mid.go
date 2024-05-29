package leetcode_go

import "sort"

func customSortString(order string, s string) string {
	m := make(map[byte]int)
	for i := 0; i < len(order); i++ {
		m[order[i]] = i + 1
	}
	t := []byte(s)
	sort.Slice(t, func(i, j int) bool { return m[t[i]] < m[t[j]] })
	return string(t)
}
