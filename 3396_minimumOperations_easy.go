package leetcode_go

import (
	"slices"
)

func minimumOperations(nums []int) int {
	cnt := map[int]bool{}
	for i, num := range slices.Backward(nums) {
		if cnt[num] {
			return i/3 + 1
		}
		cnt[num] = true
	}
	return 0
}
