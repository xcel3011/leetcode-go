package leetcode_go

// 给你一个整数数组 nums 和一个整数 k。如果某个连续子数组中恰好有 k 个奇数数字，我们就认为这个子数组是「优美子数组」。
// 请返回这个数组中 「优美子数组」 的数目。
//
// 示例 1：
// 输入：nums = [1,1,2,1,1], k = 3
// 输出：2
// 解释：包含 3 个奇数的子数组是 [1,1,2,1] 和 [1,2,1,1] 。
//
// 示例 2：
// 输入：nums = [2,4,6], k = 1
// 输出：0
// 解释：数列中不包含任何奇数，所以不存在优美子数组。
//
// 示例 3：
// 输入：nums = [2,2,2,1,2,2,1,2,2,2], k = 2
// 输出：16
//
// 提示：
// 1 <= nums.length <= 50000
// 1 <= nums[i] <= 10^5
// 1 <= k <= nums.length
func numberOfSubarrays(nums []int, k int) int {
	sum1, sum2 := 0, 0
	l1, l2, r := 0, 0, 0
	ans := 0
	for r < len(nums) {
		in := nums[r]
		r++
		if in%2 == 1 {
			sum1++
			sum2++
		}
		for sum1 > k {
			out := nums[l1]
			l1++
			if out%2 == 1 {
				sum1--
			}
		}
		for sum2 >= k {
			out := nums[l2]
			l2++
			if out%2 == 1 {
				sum2--
			}
		}
		ans += l2 - l1
	}
	return ans
}
