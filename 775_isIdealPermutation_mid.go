package leetcode_go

// 全局倒置 			nums[i]>nums[j]	i<j
// 局部倒置 			nums[i]>nums[j]	i=j-1
// 全局倒置且不是局部	nums[i]>nums[j]	i<j-1
func isIdealPermutation(nums []int) bool {
	max := 0
	for i := 2; i < len(nums); i++ {
		if nums[i-2] > max {
			max = nums[i-2]
		}
		if nums[i] > max {
			return false
		}
	}
	return true
}
