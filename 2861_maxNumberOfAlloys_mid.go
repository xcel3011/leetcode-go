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

func maxNumberOfAlloys250417(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
	left, right := 1, budget+slices.Min(stock) // 区间是(1,max+1)
	for left < right {
		mid := left + (right-left)/2
		valid := false
		for i := 0; i < k; i++ {
			spend := 0
			for j := 0; j < n; j++ {
				spend += max(composition[i][j]*mid-stock[j], 0) * cost[j]
			}
			if spend <= budget {
				valid = true
				break
			}
		}

		if valid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}
