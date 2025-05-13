package leetcode_go

import (
	"slices"
)

func minimumTime(time []int, totalTrips int) int64 {
	minT := slices.Min(time)
	l := minT - 1
	r := totalTrips * minT
	for l+1 < r { // 开区间内不能为kong
		mid := l + (r-l)/2
		sum := 0
		for _, t := range time {
			sum += mid / t
		}

		if sum >= totalTrips {
			r = mid
		} else {
			l = mid
		}
	}
	return int64(r)
}

func minimumTime250417(time []int, totalTrips int) int64 {
	mn := slices.Min(time)
	left := mn - 1             // 循环不变量：check(left) 恒为 false
	right := mn*totalTrips + 1 // 循环不变量：check(right) 恒为 true
	for left+1 < right {
		mid := left + (right-left)/2
		sum := 0
		for _, t := range time {
			sum += mid / t
		}
		if sum >= totalTrips {
			right = mid
		} else {
			left = mid
		}
	}

	// 此时 left 等于 right-1
	return int64(right)
}
