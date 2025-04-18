package leetcode_go

import "slices"

// 给你一个下标从 0 开始长度为 n 的整数数组 nums 和一个整数 target ，
// 请你返回满足 0 <= i < j < n 且 nums[i] + nums[j] < target 的下标对 (i, j) 的数目。
//
// 示例 1：
// 输入：nums = [-1,1,2,3,1], target = 2
// 输出：3
// 解释：总共有 3 个下标对满足题目描述：
// - (0, 1) ，0 < 1 且 nums[0] + nums[1] = 0 < target
// - (0, 2) ，0 < 2 且 nums[0] + nums[2] = 1 < target
// - (0, 4) ，0 < 4 且 nums[0] + nums[4] = 0 < target
// 注意 (0, 3) 不计入答案因为 nums[0] + nums[3] 不是严格小于 target 。
//
// 示例 2：
// 输入：nums = [-6,2,5,-2,-7,-1,3], target = -2
// 输出：10
// 解释：总共有 10 个下标对满足题目描述：
// - (0, 1) ，0 < 1 且 nums[0] + nums[1] = -4 < target
// - (0, 3) ，0 < 3 且 nums[0] + nums[3] = -8 < target
// - (0, 4) ，0 < 4 且 nums[0] + nums[4] = -13 < target
// - (0, 5) ，0 < 5 且 nums[0] + nums[5] = -7 < target
// - (0, 6) ，0 < 6 且 nums[0] + nums[6] = -3 < target
// - (1, 4) ，1 < 4 且 nums[1] + nums[4] = -5 < target
// - (3, 4) ，3 < 4 且 nums[3] + nums[4] = -9 < target
// - (3, 5) ，3 < 5 且 nums[3] + nums[5] = -3 < target
// - (4, 5) ，4 < 5 且 nums[4] + nums[5] = -8 < target
// - (4, 6) ，4 < 6 且 nums[4] + nums[6] = -4 < target
//
// 提示：
// 1 <= nums.length == n <= 50
// -50 <= nums[i], target <= 50
func countPairs(nums []int, target int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] < target {
				ans++
			}
		}
	}
	return ans
}

func countPairs250408(nums []int, target int) int {
	//先排序
	slices.Sort(nums)
	ans := 0
	l, r := 0, len(nums)-1
	for l < r {
		sum := nums[l] + nums[r]
		if sum < target {
			//[l,r],[l,r-1] ~ [l,l+1]都符合
			ans += r - l
			l++
		} else {
			r--
		}
	}
	return ans
}
