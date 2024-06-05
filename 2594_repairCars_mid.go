package leetcode_go

import (
	"math"
	"slices"
)

func repairCars(ranks []int, cars int) int64 {
	// 1 <= ranks.length <= 105
	// 1 <= ranks[i] <= 100
	// 1 <= cars <= 106
	l, r := 1, slices.Min(ranks)*cars*cars // 闭区间[l,r]
	for l <= r {
		mid := l + (r-l)/2
		sum := 0
		for _, rank := range ranks {
			sum += int(math.Sqrt(float64(mid / rank)))
		}

		// sum小的话说明mid小
		if sum < cars {
			l = mid + 1 //[mid+1,r]
		} else {
			r = mid - 1 //[l,mid-1]
		}
	}
	return int64(l)
}
