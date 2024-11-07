package leetcode_go

func resultsArray3255(nums []int, k int) []int {
	res := make([]int, len(nums)-k+1)
	for i := 0; i < len(res); i++ {
		res[i] = -1
	}

	cnt := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] == nums[i-1]+1 {
			cnt++
		} else {
			cnt = 1
		}
		if cnt >= k {
			res[i-k+1] = nums[i]
		}
	}
	return res
}
