package leetcode_go

import (
	"sort"
)

func minLengthAfterRemovals(nums []int) int {
	x := nums[len(nums)/2]
	count := sort.SearchInts(nums, x+1) - sort.SearchInts(nums, x)
	return max(count*2-len(nums), len(nums)%2)
}
