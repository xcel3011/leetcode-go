package leetcode_go

import (
	"math"
	"slices"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var (
		n    = len(nums)
		best = math.MaxInt32
	)

	// 根据差值的绝对值来更新答案
	update := func(cur int) {
		if abs(cur-target) < abs(best-target) {
			best = cur
		}
	}

	// 枚举 a
	for i := 0; i < n; i++ {
		// 保证和上一次枚举的元素不相等
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 使用双指针枚举 b 和 c
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			// 如果和为 target 直接返回答案
			if sum == target {
				return target
			}
			update(sum)
			if sum > target {
				// 如果和大于 target，移动 c 对应的指针
				k0 := k - 1
				// 移动到下一个不相等的元素
				for j < k0 && nums[k0] == nums[k] {
					k0--
				}
				k = k0
			} else {
				// 如果和小于 target，移动 b 对应的指针
				j0 := j + 1
				// 移动到下一个不相等的元素
				for j0 < k && nums[j0] == nums[j] {
					j0++
				}
				j = j0
			}
		}
	}
	return best
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

// 给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
// 返回这三个数的和。
// 假定每组输入只存在恰好一个解。
//
// 示例 1：
// 输入：nums = [-1,2,1,-4], target = 1
// 输出：2
// 解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2)。
//
// 示例 2：
// 输入：nums = [0,0,0], target = 1
// 输出：0
// 解释：与 target 最接近的和是 0（0 + 0 + 0 = 0）。
//
// 提示：
// 3 <= nums.length <= 1000
// -1000 <= nums[i] <= 1000
// -104 <= target <= 104
func threeSumClosest250327(nums []int, target int) int {
	_abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	slices.Sort(nums)
	ans := math.MaxInt
	for i := 0; i < len(nums); i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			// 当前和
			sum := nums[i] + nums[l] + nums[r]

			// 如果已经跟target一样就用再找
			if sum == target {
				return sum
			}

			// 找到更好的结果
			if _abs(sum-target) < _abs(ans-target) {
				ans = sum
			}

			if sum > target {
				r--
				for r > l && nums[r] == nums[r+1] {
					r--
				}
			} else {
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}
	return ans
}

func threeSumClosest250408(nums []int, target int) int {
	slices.Sort(nums)
	_abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	ans := math.MaxInt
	for i := 0; i < len(nums); i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				return target
			}
			if _abs(sum-target) < _abs(ans-target) {
				ans = sum
			}
			if sum > target {
				r--
				for r > l && nums[r] == nums[r+1] {
					r--
				}
			} else if sum < target {
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}
	return ans
}
