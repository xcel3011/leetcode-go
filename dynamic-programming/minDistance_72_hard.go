package dynamic_programming

var minDistanceMemo [][]int

func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	minDistanceMemo = make([][]int, n)
	for i := 0; i < n; i++ {
		minDistanceMemo[i] = make([]int, m)
		for j := 0; j < m; j++ {
			minDistanceMemo[i][j] = -1
		}
	}
	return minDistanceDP(word1, word2, n-1, m-1)
}

func minDistanceDP(s1, s2 string, n, m int) int {
	if n == -1 {
		return 0
	}
	if m == -1 {
		return 0
	}
	if minDistanceMemo[n][m] != -1 {
		return minDistanceMemo[n][m]
	}
	if s1[n] == s2[m] {
		minDistanceMemo[n][m] = minDistanceDP(s1, s2, n-1, m-1)
	}
	minDistanceMemo[n][m] = min(
		minDistanceDP(s1, s2, n-1, m)+1,   // 删除
		minDistanceDP(s1, s2, n-1, m-1)+1, // 替换
		minDistanceDP(s1, s2, n, m-1)+1,   // 插入
	)
	return minDistanceMemo[n][m]
}
