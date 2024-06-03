package leetcode_go

import (
	"sort"
)

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	var ans int64
	for i := 0; i < len(nums); i++ {
		r := sort.SearchInts(nums[:i], upper-nums[i]+1)
		l := sort.SearchInts(nums[:i], nums[i]-lower)
		ans += int64(r - l)
	}
	return ans
}
