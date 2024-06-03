package leetcode_go

import (
	"sort"
)

func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)

	// 求前缀和
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	for i := 0; i < len(queries); i++ {
		queries[i] = sort.SearchInts(nums, queries[i]+1)
	}
	return queries
}
