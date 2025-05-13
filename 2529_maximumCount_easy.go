package leetcode_go

import (
	"slices"
)

func maximumCount(nums []int) int {
	f := func(nums []int, target int) int {
		left, right := 0, len(nums)
		for left < right {
			mid := left + (right-left)/2
			if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid
			}
		}
		return left
	}

	return max(f(nums, 0), len(nums)-f(nums, 1))
}

func maximumCount250417(nums []int) int {
	// 找到第一个大于等于0的位置i,从[0, i-1]是负数
	neg, _ := slices.BinarySearch(nums, 0)
	// 找到第一个大于1的位置j,从[j, len(nums)-1]是正数
	pos, _ := slices.BinarySearch(nums, 1)
	return max(neg, len(nums)-pos)
}
