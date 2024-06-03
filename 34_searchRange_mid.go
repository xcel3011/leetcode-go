package leetcode_go

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
