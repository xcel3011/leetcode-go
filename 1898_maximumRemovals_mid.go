package leetcode_go

import (
	"sort"
)

func maximumRemovals(s string, p string, removable []int) int {
	return sort.Search(len(removable), func(k int) bool {
		del := make([]bool, len(s))
		for _, idx := range removable[:k+1] {
			del[idx] = true
		}
		j := 0
		for i := range s {
			if !del[i] && s[i] == p[j] {
				j++
				if j == len(p) {
					return false
				}
			}
		}
		return true
	})
}
