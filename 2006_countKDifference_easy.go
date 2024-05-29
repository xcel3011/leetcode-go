package leetcode_go

func countKDifference(nums []int, k int) int {
	var count int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]-nums[j] == k || nums[j]-nums[i] == k {
				count++
			}
		}
	}
	return count
}
