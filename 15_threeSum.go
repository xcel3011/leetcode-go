package leetcode_go

import (
	"slices"
	"sort"
)

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		lo, hi := i+1, len(nums)-1
		for lo < hi {
			sum := nums[i] + nums[lo] + nums[hi]
			if sum > 0 {
				hi--
			} else if sum < 0 {
				lo++
			} else {
				res = append(res, []int{nums[i], nums[lo], nums[hi]})
				for lo < hi && nums[lo] == nums[lo+1] {
					lo++
				}
				for lo < hi && nums[hi] == nums[hi-1] {
					hi--
				}
				lo++
				hi--
			}
		}
	}
	return res
}

// 给你一个整数数组 nums
// 判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k
// 同时还满足 nums[i] + nums[j] + nums[k] == 0
// 请你返回所有和为 0 且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
// 示例 1：
// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]
// 解释：
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
// 不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
// 注意，输出的顺序和三元组的顺序并不重要。
//
// 示例 2：
// 输入：nums = [0,1,1]
// 输出：[]
// 解释：唯一可能的三元组和不为 0 。
//
// 示例 3：
// 输入：nums = [0,0,0]
// 输出：[[0,0,0]]
// 解释：唯一可能的三元组和为 0 。
//
// 提示：
// 3 <= nums.length <= 3000
// -105 <= nums[i] <= 105
func threeSum250327(nums []int) (res [][]int) {
	if len(nums) < 3 {
		return
	}

	// 先排序
	slices.Sort(nums)
	for i := 0; i < len(nums); i++ {
		// 排序后第一个数大于0，直接返回
		if nums[i] > 0 {
			return
		}

		// 保证第一位和之前已经存在的答案不一样，防止重复
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 从i+1开始双指针往中间走
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})

				// 保证第二位和第三位和之前已经存在的答案不一样，防止重复
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}

				// 继续往中间走
				l++
				r--
			}
		}
	}
	return
}

func threeSum250408(nums []int) [][]int {
	//先从小到大排序
	slices.Sort(nums)

	var ans [][]int

	//枚举I
	for i, x := range nums[:len(nums)-2] {
		//跳过重复
		if i > 0 && x == nums[i-1] {
			continue
		}
		//最小的三个都大于零
		if x+nums[i+1]+nums[i+2] > 0 {
			break
		}
		//x加最大的两个小于0
		if x+nums[len(nums)-1]+nums[len(nums)-2] < 0 {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			sum := x + nums[l] + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				ans = append(ans, []int{x, nums[l], nums[r]})
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				r--
				for r > l && nums[r] == nums[r+1] {
					r--
				}
			}

		}
	}
	return ans
}
