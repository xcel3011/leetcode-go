package dynamic_programming

// O(N^2)
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
