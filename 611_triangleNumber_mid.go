package leetcode_go

import (
	"slices"
)

// 给定一个包含非负整数的数组 nums ，返回其中可以组成三角形三条边的三元组个数。
//
// 示例 1:
// 输入: nums = [2,2,3,4]
// 输出: 3
// 解释:有效的组合是:
// 2,3,4 (使用第一个 2)
// 2,3,4 (使用第二个 2)
// 2,2,3
//
// 示例 2:
// 输入: nums = [4,2,3,4]
// 输出: 4
//
// 提示:
// 1 <= nums.length <= 1000
// 0 <= nums[i] <= 1000
func triangleNumber(nums []int) int {
	slices.Sort(nums)
	ans := 0
	for i := 2; i < len(nums); i++ {
		l, r := 0, i-1
		for l < r {
			sum := nums[l] + nums[r]
			if sum > nums[i] {
				ans += r - l
				r--
			} else {
				l++
			}
		}
	}
	return ans
}

func triangleNumber250408(nums []int) int {
	slices.Sort(nums)
	ans := 0
	for i := 2; i < len(nums); i++ {
		l, r := 0, i-1
		for l < r {
			sum := nums[l] + nums[r]
			if sum > nums[i] {
				ans += r - l
				r--
			} else {
				l++
			}
		}
	}
	return ans

}
