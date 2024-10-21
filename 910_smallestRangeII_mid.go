package leetcode_go

import (
	"sort"
)

func smallestRangeII(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	diff := nums[n-1] - nums[0]
	for i := 1; i < n; i++ {
		a := max(nums[i-1]+k, nums[n-1]-k)
		b := min(nums[i]-k, nums[0]+k)
		diff = min(diff, a-b)
	}
	return diff
}
