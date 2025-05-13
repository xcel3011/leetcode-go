package leetcode_go

import (
	"slices"
)

func searchRange(nums []int, target int) []int {
	f := func(nums []int, target int) int {
		left, right := 0, len(nums) // 左闭右开区间 [left, right)
		for left < right {          // 区间不为空
			mid := left + (right-left)/2
			// nums[left-1] < target
			// nums[right] >= target
			if nums[mid] < target {
				left = mid + 1 // 范围缩小到 [mid+1, right)
			} else {
				right = mid // 范围缩小到 [left, mid)
			}
		}
		return left
	}

	start := f(nums, target)
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	// 如果 start 存在，那么 end 必定存在
	end := f(nums, target+1) - 1
	return []int{start, end}
}

// 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
//
// 如果数组中不存在目标值 target，返回 [-1, -1]。
//
// 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
func searchRange250321(nums []int, target int) []int {
	start, ok := slices.BinarySearch(nums, target)
	if !ok {
		return []int{-1, -1}
	}
	end, _ := slices.BinarySearch(nums, target+1)
	return []int{start, end - 1}
}

func searchRange250417(nums []int, target int) []int {
	begin, ok := slices.BinarySearch(nums, target)
	if !ok {
		return []int{-1, -1}
	}
	end, _ := slices.BinarySearch(nums, target+1)
	return []int{begin, end - 1}
}
