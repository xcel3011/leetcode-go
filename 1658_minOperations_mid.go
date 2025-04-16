package leetcode_go

import (
	"math"
)

// 给你一个整数数组 nums 和一个整数 x 。
// 每一次操作时，你应当移除数组 nums 最左边或最右边的元素，然后从 x 中减去该元素的值。
// 请注意，需要 修改 数组以供接下来的操作使用。
// 如果可以将 x 恰好 减到 0 ，返回 最小操作数 ；否则，返回 -1 。
//
// 示例 1：
// 输入：nums = [1,1,4,2,3], x = 5
// 输出：2
// 解释：最佳解决方案是移除后两个元素，将 x 减到 0 。
//
// 示例 2：
// 输入：nums = [5,6,7,8,9], x = 4
// 输出：-1
//
// 示例 3：
// 输入：nums = [3,2,20,1,1,3], x = 10
// 输出：5
// 解释：最佳解决方案是移除后三个元素和前两个元素（总共 5 次操作），将 x 减到 0 。
//
// 提示：
// 1 <= nums.length <= 105
// 1 <= nums[i] <= 104
// 1 <= x <= 109
func minOperations250325(nums []int, x int) int {
	// 先计算总和
	sum := 0
	for _, v := range nums {
		sum += v
	}
	// 中间滑动的和
	target := sum - x
	if target < 0 {
		return -1
	}

	// 滑动窗口
	l, r := 0, 0
	cnt := 0
	ans := math.MaxInt

	for r < len(nums) {
		in := nums[r]
		r++
		cnt += in
		for cnt > target {
			out := nums[l]
			l++
			cnt -= out
		}
		if cnt == target {
			ans = min(ans, len(nums)-(r-l))
		}
	}

	if ans == math.MaxInt {
		return -1
	}

	return ans
}

func minOperations250416(nums []int, x int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	target := sum - x
	if target < 0 {
		return -1
	}
	ans := math.MaxInt
	cnt := 0
	left := 0
	for right, num := range nums {
		cnt += num
		for cnt > target {
			cnt -= nums[left]
			left++
		}
		if cnt == target {
			ans = min(ans, len(nums)-(right-left+1))
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
