package leetcode_go

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
