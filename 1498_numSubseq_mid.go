package leetcode_go

import (
	"slices"
)

func numSubseq(nums []int, target int) int {
	slices.Sort(nums)
	ans := 0
	l, r := 0, len(nums)-1

	// 如果 nums[left] + nums[right] ≤ target
	// 说明 [left, right] 之间的所有子序列都满足条件
	// 共有 2^(right - left) 个（因为 right - left 个元素可独立选择是否加入）
	pow := map[int]int{0: 1}
	for i := 1; i < len(nums); i++ {
		pow[i] = pow[i-1] * 2 % 1000000007
	}

	for l <= r {
		if nums[l]+nums[r] > target {
			r--
		} else {
			ans += pow[r-l]
			l++
		}
	}
	return ans % 1000000007
}
