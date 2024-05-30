package leetcode_go

import (
	"slices"
)

func minimumDifference(nums []int, k int) int {
	slices.SortStableFunc(nums, func(a, b int) int {
		return a - b
	})

	ans := nums[k-1] - nums[0]
	for i := k; i < len(nums); i++ {
		ans = min(ans, nums[i]-nums[i-k+1])
	}
	return ans
}
