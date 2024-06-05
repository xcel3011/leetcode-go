package leetcode_go

import (
	"math"
	"sort"
)

func smallestDivisor(nums []int, threshold int) int {
	sort.Ints(nums)
	ans := -1
	l, r := 1, nums[len(nums)-1]
	for l <= r {
		mid := l + (r-l)/2
		total := 0
		for _, num := range nums {
			total += int(math.Ceil(float64(num) / float64(mid)))
		}

		if total <= threshold {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}
