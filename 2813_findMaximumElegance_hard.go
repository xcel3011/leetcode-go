package leetcode_go

import (
	"slices"
)

func findMaximumElegance(items [][]int, k int) int64 {
	// 按利润从大到小排序
	slices.SortStableFunc(items, func(a, b []int) int {
		return b[0] - a[0]
	})

	var (
		ans, totalProfit int
		duplicate        []int // 记录重复出现类别的利润
	)
	visited := map[int]bool{} // 已经遍历过的类别
	for i, item := range items {
		profit, category := item[0], item[1]
		if i < k { // 前k个直接统计
			// 统计利润
			totalProfit += profit
			if !visited[category] { // 第一次出现该类别只做标记
				visited[category] = true
			} else { // 后续再出现该类别，将利润入栈
				duplicate = append(duplicate, profit)
			}
		} else if !visited[category] && len(duplicate) > 0 { // 该类别没出现过，之前统计的类别有重复出现
			visited[category] = true
			totalProfit += profit - duplicate[len(duplicate)-1]
			duplicate = duplicate[:len(duplicate)-1]
		}
		ans = max(ans, totalProfit+len(visited)*len(visited))
	}
	return int64(ans)
}
