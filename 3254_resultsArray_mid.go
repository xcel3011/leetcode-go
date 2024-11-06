package leetcode_go

func resultsArray(nums []int, k int) []int {
	ans := make([]int, len(nums)-k+1)
	for i := 0; i < len(ans); i++ {
		ans[i] = -1
	}

	var cnt int
	for i := 0; i < len(nums); i++ {

		if i == 0 || nums[i] == nums[i-1]+1 {
			cnt++
		} else {
			cnt = 1
		}
		if cnt >= k {
			ans[i-k+1] = nums[i]
		}
	}
	return ans
}
