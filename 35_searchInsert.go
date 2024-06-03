package leetcode_go

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
