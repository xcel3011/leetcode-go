package leetcode_go

import (
	"slices"
)

func maximumLength2group(s string) int {
	groups := [26][]int{}
	count := 0
	for i := 0; i < len(s); i++ {
		count++
		if i+1 == len(s) || s[i] != s[i+1] {
			groups[s[i]-'a'] = append(groups[s[i]-'a'], count)
			count = 0
		}
	}

	ans := 0
	for i := 0; i < len(groups); i++ {
		group := groups[i]
		if len(group) == 0 {
			continue
		}
		slices.SortStableFunc(group, func(a, b int) int {
			return b - a
		})
		group = append(group, 0, 0)
		ans = max(ans, group[0]-2, min(group[0]-1, group[1]), group[2])
	}
	if ans == 0 {
		return -1
	}
	return ans
}
