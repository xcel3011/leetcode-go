package leetcode_go

import (
	"slices"
)

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		}
	}
	return left
}

func searchInsert2420603(nums []int, target int) int {
	left, right := 0, len(nums) // 开区间[0,n)
	for left < right {          // 区间不为空
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1 // 区间变为[mid+1,right)
		} else {
			right = mid // 区间变为[left,mid)
		}
	}
	return left
}

// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
// 如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
// 请必须使用时间复杂度为 O(log n) 的算法。
func searchInsert250321(nums []int, target int) int {
	idx, _ := slices.BinarySearch(nums, target)
	return idx
}
