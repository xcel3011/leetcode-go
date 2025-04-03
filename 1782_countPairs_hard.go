package leetcode_go

import (
	"sort"
)

func countPairs250331(n int, edges [][]int, queries []int) []int {
	deg := make([]int, n+1)
	type pair struct{ x, y int }
	cnt := map[pair]int{}
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		if x > y {
			x, y = y, x
		}

		deg[x]++
		deg[y]++
		// 统计每条边的出现次数，注意 1-2 和 2-1 算同一条边
		cnt[pair{x, y}]++
	}

	ans := make([]int, len(queries))
	// 统计度数之和大于 q 的边的数量
	// 所以需要减去 cnt[pair{x, y}]
	sortedDeg := append([]int(nil), deg...)
	sort.Ints(sortedDeg)
	for i, q := range queries {
		l, r := 1, n
		for l < r {
			if sortedDeg[l]+sortedDeg[r] <= q {
				l++
			} else {
				ans[i] += r - l
				r--
			}
		}

		// 注意 1-2 和 2-1 算同一条边
		for p, c := range cnt {
			sum := deg[p.x] + deg[p.y]
			if sum > q && sum-c <= q {
				ans[i]--
			}
		}
	}
	return ans
}
