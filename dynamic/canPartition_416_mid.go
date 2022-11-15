package dynamic

func canPartition(nums []int) bool {
	var sum int
	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}

	dp := make([][]bool, len(nums)+1)
	for i := range dp {
		dp[i] = make([]bool, sum/2+1)
		dp[i][0] = true
	}

	for i := 1; i <= len(nums); i++ {
		for j := 1; j <= sum/2; j++ {
			if j-nums[i-1] < 0 {
				// 背包容量不足，不能装入第 i 个物品
				dp[i][j] = dp[i-1][j]
			} else {
				// 装入或不装入背包
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[len(nums)][sum/2]
}
