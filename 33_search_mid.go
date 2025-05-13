package leetcode_go

func search250213(nums []int, target int) int {
	n := len(nums)
	end := nums[n-1]
	check := func(i int) bool {
		if nums[i] >= end {
			return target > end && target <= nums[i]
		} else {
			return target > end || target <= nums[i]
		}
	}

	left, right := -1, n-1
	for left+1 < right {
		mid := left + (right-left)/2
		if check(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	// fmt.Println(right, nums[right], target)
	if nums[right] != target {
		return -1
	}
	return right
}
