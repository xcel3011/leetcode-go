package leetcode_go

import (
	"slices"
)

func minEatingSpeed(piles []int, h int) int {
	l, r := 1, slices.Max(piles) // 区间是[1,max]
	for l <= r {
		mid := l + (r-l)/2
		total := 0
		for _, p := range piles {
			total += (p-1)/mid + 1 // 向上取整
		}
		if total > h { // 耗时大说明mid小了,所以要向右移动
			l = mid + 1 //[mid+1,r]
		} else {
			r = mid - 1 //[l,mid+1]
		}
	}
	return l
}
