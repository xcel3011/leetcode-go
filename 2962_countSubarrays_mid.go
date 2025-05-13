package leetcode_go

import (
	"slices"
)

// 给你一个整数数组 nums 和一个 正整数 k 。
// 请你统计有多少满足 「 nums 中的 最大 元素」至少出现 k 次的子数组，并返回满足这一条件的子数组的数目。
// 子数组是数组中的一个连续元素序列。
//
// 示例 1：
// 输入：nums = [1,3,2,3,3], k = 2
// 输出：6
// 解释：包含元素 3 至少 2 次的子数组为：[1,3,2,3]、[1,3,2,3,3]、[3,2,3]、[3,2,3,3]、[2,3,3] 和 [3,3] 。
//
// 示例 2：
// 输入：nums = [1,4,2,1], k = 3
// 输出：0
// 解释：没有子数组包含元素 4 至少 3 次。
//
// 提示：
// 1 <= nums.length <= 105
// 1 <= nums[i] <= 106
// 1 <= k <= 105
func countSubarrays(nums []int, k int) int64 {
	mx := slices.Max(nums)
	cnt := 0
	l, r := 0, 0
	ans := 0
	for r < len(nums) {
		in := nums[r]
		r++
		if in == mx {
			cnt++
		}
		for cnt == k {
			out := nums[l]
			l++
			if out == mx {
				cnt--
			}
		}
		ans += l
	}
	return int64(ans)
}

func countSubarrays250429mid(nums []int, k int) int64 {
	mx := slices.Max(nums)
	cnt := 0
	left, right := 0, 0
	ans := 0
	for right < len(nums) {
		// 判断右侧进入窗口是否为最大值
		if nums[right] == mx {
			cnt++
		}

		// 右指针移动
		right++

		// 如果当前满足k
		for cnt == k {
			//
			if nums[left] == mx {
				cnt--
			}
			left++
		}

		//right=4,left=2
		//[3,2,3],[1,3,2,3]
		//right=5,left=5
		//[3,3],[2,3,3],[3,2,3,3],[1,3,2,3,3]
		ans += left
	}
	return int64(ans)
}

func countSubarrays250416(nums []int, k int) int64 {
	ans := 0
	cnt := 0
	mx := slices.Max(nums)
	left := 0
	for _, x := range nums {
		if x == mx {
			cnt++
		}
		for cnt == k {
			if nums[left] == mx {
				cnt--
			}
			left++
		}
		ans += left
	}
	return int64(ans)
}
