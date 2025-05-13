package leetcode_go

func findMin(nums []int) int {
	//[0,n-1)
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[right] { // 蓝色
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}
