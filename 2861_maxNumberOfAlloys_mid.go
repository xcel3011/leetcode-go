package leetcode_go

import (
	"slices"
)

func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
	l, r := 1, budget+slices.Min(stock)
	for l <= r {
		mid := l + (r-l)/2
		valid := false
		for i := 0; i < k; i++ {
			spend := 0
			for j := 0; j < n; j++ {
				spend += max(composition[i][j]*mid-stock[j], 0) * cost[j]
			}
			if spend <= budget {
				valid = true
			}
		}
		if valid {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r
}
