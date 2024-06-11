package leetcode_go

func maxCoins(nums []int) int {
	n := len(nums)
	nums = append(append([]int{1}, nums...), 1)

	// 状态转换方程
	// f[i][j]=max(f[i][j],f[i][k]+nums[i]*nums[k]*nums[j]+f[k][j])
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j < n+2; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+nums[i]*nums[k]*nums[j])
			}
		}
	}
	return dp[0][n+1]
}
