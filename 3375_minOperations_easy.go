package leetcode_go

import (
	"slices"
)

func minOperations250409(nums []int, k int) int {
	mn := slices.Min(nums)
	if k > mn {
		return -1
	}

	cnt := map[int]int{}
	for _, num := range nums {
		if num >= k {
			cnt[num]++
		}
	}
	if k == mn {
		return len(cnt) - 1
	}
	return len(cnt)
}
