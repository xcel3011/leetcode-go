package leetcode_go

import (
	"slices"
)

func minDays(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}
	//[l,r]
	l, r := slices.Min(bloomDay), slices.Max(bloomDay)
	for l <= r {
		mid := l + (r-l)/2
		flowers, bouquets := 0, 0
		for _, d := range bloomDay {
			if d > mid { // 成长需要的时间大于当前的天数
				flowers = 0
			} else {
				flowers++
				if flowers == k { // 连续的花满足k
					flowers = 0
					bouquets++
				}
			}
		}
		if bouquets < m {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
