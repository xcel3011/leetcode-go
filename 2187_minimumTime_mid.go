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
