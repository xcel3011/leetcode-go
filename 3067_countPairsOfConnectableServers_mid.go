package leetcode_go

func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	type edge struct{ to, wt int }
	n := len(edges) + 1
	g := make([][]edge, n)
	for i := 0; i < len(edges); i++ {
		x, y, wt := edges[i][0], edges[i][1], edges[i][2]
		g[x] = append(g[x], edge{y, wt})
		g[y] = append(g[y], edge{x, wt})
	}
	ans := make([]int, n)
	for i, gi := range g {
		var cnt int
		var dfs func(x, fa, sum int)
		dfs = func(x, fa, sum int) {
			if sum%signalSpeed == 0 {
				cnt++
			}
			for _, e := range g[x] {
				if e.to != fa {
					dfs(e.to, x, sum+e.wt)
				}
			}
			return
		}
		sum := 0
		for _, e := range gi {
			cnt = 0
			dfs(e.to, i, e.wt)
			ans[i] += cnt * sum
			sum += cnt
		}
	}
	return ans
}
