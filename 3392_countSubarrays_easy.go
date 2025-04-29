package leetcode_go

func countSubarrays250429(nums []int) int {
	ans := 0
	for i := 0; i < len(nums)-2; i++ {
		if float64(nums[i]+nums[i+2]) == float64(nums[i+1])/2 {
			ans++
		}
	}
	return ans
}
