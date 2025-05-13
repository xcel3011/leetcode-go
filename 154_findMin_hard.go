package leetcode_go

import (
	"fmt"
)

func findMin250513(nums []int) int {
	left, right := -1, len(nums)-1
	for left+1 < right {
		// 防止溢出
		mid := left + (right-left)/2
		fmt.Println(mid, nums[mid], nums[right])
		if nums[mid] == nums[right] {
			right--
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid
		}
	}
	return nums[right]
}
