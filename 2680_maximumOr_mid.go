package leetcode_go

// 给你一个下标从 0 开始长度为 n 的整数数组 nums 和一个整数 k 。每一次操作中，你可以选择一个数并将它乘 2 。
//
// 你最多可以进行 k 次操作，请你返回 nums[0] | nums[1] | ... | nums[n - 1] 的最大值。
//
// a | b 表示两个整数 a 和 b 的 按位或 运算。
//
// 示例 1：
//
// 输入：nums = [12,9], k = 1
// 输出：30
// 解释：如果我们对下标为 1 的元素进行操作，新的数组为 [12,18] 。此时得到最优答案为 12 和 18 的按位或运算的结果，也就是 30 。
// 示例 2：
//
// 输入：nums = [8,1,2], k = 2
// 输出：35
// 解释：如果我们对下标 0 处的元素进行操作，得到新数组 [32,1,2] 。此时得到最优答案为 32|1|2 = 35 。
func maximumOr(nums []int, k int) int64 {
	n := len(nums)
	suf := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		// suf[i] 表示 nums[i+1] 到 nums[n-1] 的 OR
		suf[i] = nums[i+1] | suf[i+1]
	}

	// pre 表示 nums[0] 到 nums[i-1] 的 OR
	ans, pre := 0, 0
	for i := 0; i < len(nums); i++ {
		ans = max(ans, pre|nums[i]<<k|suf[i])
		pre |= nums[i]
	}
	return int64(ans)
}
