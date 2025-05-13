package leetcode_go

func findPeakElement(nums []int) int {
	//[0,n-2]
	//(-1,n-1)
	left, right := -1, len(nums)-1
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] { // 蓝色
			right = mid
		} else {
			left = mid
		}
	}
	return right
}
