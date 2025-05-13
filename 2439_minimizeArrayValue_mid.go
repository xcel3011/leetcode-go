package leetcode_go

import (
	"slices"
	"sort"
)

func minimizeArrayValue(nums []int) int {
	return sort.Search(slices.Max(nums), func(limit int) bool {
		// 计算nums中每个元素减去limit后的和
		extra := 0
		for i := len(nums) - 1; i >= 0; i-- {
			extra = max(nums[i]+extra-limit, 0)
		}
		return nums[0]+extra <= limit
	})
}
