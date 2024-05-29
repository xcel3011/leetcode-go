package leetcode_go

import "math"

var memo [][]int

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	res := math.MaxInt
	memo = make([][]int, n)

	// 备忘录里的值初始化为 66666
	for i := 0; i < n; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = 66666
		}
	}
	// 终点可能在 matrix[n-1] 的任意一列
	for i := 0; i < n; i++ {
		res = min(res, dp(matrix, n-1, i))
	}
	return res
}

func dp(matrix [][]int, i, j int) int {
	// 1. 索引合法性检查
	if i < 0 || j < 0 || i >= len(matrix) || j >= len(matrix[0]) {
		return 99999
	}

	// 2. base case
	if i == 0 {
		return matrix[i][j]
	}

	// 3. 查找备忘录，防止重复计算
	if memo[i][j] != 66666 {
		return memo[i][j]
	}

	// 4. 进行状态转移
	memo[i][j] = matrix[i][j] + min(
		dp(matrix, i-1, j-1),
		dp(matrix, i-1, j),
		dp(matrix, i-1, j+1),
	)

	return memo[i][j]
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v < res {
			res = v
		}
	}
	return res
}
