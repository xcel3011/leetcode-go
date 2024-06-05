package leetcode_go

import (
	"sort"
)

func shipWithinDays(weights []int, days int) int {
	// 左边界是所有重量的最大值
	// 右边界是重量和
	l, r := 0, 0
	for _, w := range weights {
		l = max(l, w)
		r += w
	}
	return l + sort.Search(r-l, func(x int) bool {
		x += l
		day := 1
		sum := 0
		for _, w := range weights {
			if sum+w > x {
				day++
				sum = 0
			}
			sum += w
		}
		return day <= days
	})
}
