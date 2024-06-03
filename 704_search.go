package leetcode_go

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1 // 区间为[0,n-1]

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid // 找到
		} else if nums[mid] > target {
			right = mid - 1 // 区间更新为[left,mid-1]
		} else if nums[mid] < target {
			left = mid + 1 // 区间更新为[mi+1,right]
		}
	}
	return -1
}

func search1(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] < target:
			left = mid + 1
		case nums[mid] > target:
			right = mid
		}
	}
	return -1
}
