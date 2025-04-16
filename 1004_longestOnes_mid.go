package leetcode_go

// 给定一个二进制数组 nums 和一个整数 k，
// 假设最多可以翻转 k 个 0 ，则返回执行操作后 数组中连续 1 的最大个数 。
//
// 示例 1：
// 输入：nums = [1,1,1,0,0,0,1,1,1,1,0], K = 2
// 输出：6
// 解释：[1,1,1,0,0,1,1,1,1,1,1]
// 粗体数字从 0 翻转到 1，最长的子数组长度为 6。
//
// 示例 2：
// 输入：nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
// 输出：10
// 解释：[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
// 粗体数字从 0 翻转到 1，最长的子数组长度为 10。
//
// 提示：
// 1 <= nums.length <= 105
// nums[i] 不是 0 就是 1
// 0 <= k <= nums.length
func longestOnes(nums []int, k int) int {
	ans := 0
	l, r := 0, 0
	cnt := 0
	for r < len(nums) {
		in := nums[r]
		r++
		if in == 0 {
			cnt++
		}

		for cnt > k {
			out := nums[l]
			l++
			if out == 0 {
				cnt--
			}
		}
		ans = max(ans, r-l)
	}
	return ans
}

func longestOnes250416(nums []int, k int) int {
	ans := 0
	cnt := 0
	left := 0
	for right, x := range nums {
		if x == 0 {
			cnt++
		}
		for cnt > k {
			if nums[left] == 0 {
				cnt--
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
