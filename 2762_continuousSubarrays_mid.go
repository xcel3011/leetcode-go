package leetcode_go

// 给你一个下标从 0 开始的整数数组 nums 。nums 的一个子数组如果满足以下条件，那么它是 不间断 的：
// i，i + 1 ，...，j  表示子数组中的下标。对于所有满足 i <= i1, i2 <= j 的下标对，都有 0 <= |nums[i1] - nums[i2]| <= 2 。
// 请你返回 不间断 子数组的总数目。
// 子数组是一个数组中一段连续 非空 的元素序列。
//
// 示例 1：
// 输入：nums = [5,4,2,4]
// 输出：8
// 解释：
// 大小为 1 的不间断子数组：[5], [4], [2], [4] 。
// 大小为 2 的不间断子数组：[5,4], [4,2], [2,4] 。
// 大小为 3 的不间断子数组：[4,2,4] 。
// 没有大小为 4 的不间断子数组。
// 不间断子数组的总数目为 4 + 3 + 1 = 8 。
// 除了这些以外，没有别的不间断子数组。
//
// 示例 2：
// 输入：nums = [1,2,3]
// 输出：6
// 解释：
// 大小为 1 的不间断子数组：[1], [2], [3] 。
// 大小为 2 的不间断子数组：[1,2], [2,3] 。
// 大小为 3 的不间断子数组：[1,2,3] 。
// 不间断子数组的总数目为 3 + 2 + 1 = 6 。
//
// 提示：
// 1 <= nums.length <= 105
// 1 <= nums[i] <= 109
func continuousSubarrays(nums []int) int64 {
	cnt := map[int]int{}
	l, r := 0, 0
	ans := 0
	for r < len(nums) {
		in := nums[r]
		r++
		cnt[in]++
		for {
			mx, mn := in, in
			for k := range cnt {
				mx = max(mx, k)
				mn = min(mn, k)
			}
			if mx-mn <= 2 {
				break
			}

			out := nums[l]
			l++
			cnt[out]--
			if cnt[out] == 0 {
				delete(cnt, out)
			}
		}
		ans += r - l
	}
	return int64(ans)
}
