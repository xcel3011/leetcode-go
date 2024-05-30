package leetcode_go

// 动态规划
func getMoneyAmount(n int) int {
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := n - 1; i >= 1; i-- {
		for j := i + 1; j <= n; j++ {
			f[i][j] = j + f[i][j-1]
			for k := i; k < j; k++ {
				cost := k + max(f[i][k-1], f[k+1][j])
				if cost < f[i][j] {
					f[i][j] = cost
				}
			}
		}
	}
	return f[1][n]
}
