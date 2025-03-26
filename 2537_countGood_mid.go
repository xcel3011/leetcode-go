package leetcode_go

// 给你一个整数数组 nums 和一个整数 k ，请你返回 nums 中 好 子数组的数目。
// 一个子数组 arr 如果有 至少 k 对下标 (i, j) 满足 i < j 且 arr[i] == arr[j] ，那么称它是一个 好 子数组。
// 子数组 是原数组中一段连续 非空 的元素序列。
//
// 示例 1：
// 输入：nums = [1,1,1,1,1], k = 10
// 输出：1
// 解释：唯一的好子数组是这个数组本身。
//
// 示例 2：
// 输入：nums = [3,1,4,3,2,2,4], k = 2
// 输出：4
// 解释：总共有 4 个不同的好子数组：
// - [3,1,4,3,2,2] 有 2 对。
// - [3,1,4,3,2,2,4] 有 3 对。
// - [1,4,3,2,2,4] 有 2 对。
// - [4,3,2,2,4] 有 2 对。
//
// 提示：
// 1 <= nums.length <= 105
// 1 <= nums[i], k <= 109
func countGood(nums []int, k int) int64 {
	cnt := map[int]int{}
	ans := 0
	pairs := 0
	l, r := 0, 0
	for r < len(nums) {
		// 如果当前窗口中已经有c个元素x了，再来一个x，会增加c对
		in := nums[r]
		r++
		pairs += cnt[in]
		cnt[in]++
		for pairs >= k {
			out := nums[l]
			l++
			// 如果当前窗口中已经有c个元素x了，去掉一个x，会减少c-1对
			pairs -= cnt[out] - 1
			cnt[out]--
		}
		ans += l

	}
	return int64(ans)
}
