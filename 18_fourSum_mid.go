package leetcode_go

import (
	"fmt"
	"slices"
)

// 给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。
// 请你找出并返回满足下述全部条件且不重复的四元组
// [nums[a], nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
// 0 <= a, b, c, d < n
// a、b、c 和 d 互不相同
// nums[a] + nums[b] + nums[c] + nums[d] == target
// 你可以按 任意顺序 返回答案 。
//
// 示例 1：
// 输入：nums = [1,0,-1,0,-2,2], target = 0
// 输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
//
// 示例 2：
// 输入：nums = [2,2,2,2,2], target = 8
// 输出：[[2,2,2,2]]
//
// 提示：
// 1 <= nums.length <= 200
// -109 <= nums[i] <= 109
// -109 <= target <= 109
func fourSum(nums []int, target int) (ans [][]int) {
	// 排序
	slices.Sort(nums)
	for i := 0; i < len(nums)-3; i++ {
		// 去除重复情况
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 第一个数超过target不进行后续循环
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		} else if nums[i]+nums[len(nums)-1]+nums[len(nums)-2]+nums[len(nums)-3] < target {
			continue
		}

		// 从i+1开始进行三数之和
		for j := i + 1; j < len(nums)-2; j++ {
			// 去除重复情况
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			// 超过target不进行后续循环
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			} else if nums[i]+nums[j]+nums[len(nums)-1]+nums[len(nums)-2] < target {
				continue
			}

			// 从j+1开始进行两数之和
			l, r := j+1, len(nums)-1
			fmt.Println(i, j, l, r)
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum > target {
					r--
				} else if sum < target {
					l++
				} else {
					ans = append(ans, []int{nums[i], nums[j], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				}
			}
		}
	}
	return
}
