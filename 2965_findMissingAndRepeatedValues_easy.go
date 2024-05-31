package leetcode_go

func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	m := make([]int, n*n+1)

	for i := 0; i < n; i++ {
		for j := 0; j < len(grid[i]); j++ {
			m[grid[i][j]]++
		}
	}

	ans := make([]int, 2)
	for i := 1; i <= n*n; i++ {
		if m[i] == 0 {
			ans[1] = i
		} else if m[i] == 2 {
			ans[0] = i
		}
	}
	return ans
}
