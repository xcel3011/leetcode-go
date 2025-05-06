package leetcode_go

import (
	"slices"
)

func largestDivisibleSubset(nums []int) []int {
	// 先排序
	slices.Sort(nums)
	n := len(nums)
	// 初始化dp数组
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	// 找出最大子集长度和最大子集的元素
	maxSize, maxVal := 1, 1
	for i := 1; i < n; i++ {
		for j, v := range nums[:i] {
			if nums[i]%v == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxSize {
			maxSize, maxVal = dp[i], nums[i]
		}
	}

	// 如果最大子集长度为1，返回单独的元素
	if maxSize == 1 {
		return []int{nums[0]}
	}

	var ans []int
	for i := n - 1; i >= 0 && maxSize >= 0; i-- {
		if dp[i] == maxSize && maxVal%nums[i] == 0 {
			ans = append(ans, nums[i])
			maxSize--
			maxVal = nums[i]
		}
	}
	return ans
}
