package leetcode_go

// 给你一个长桌子，桌子上盘子和蜡烛排成一列。给你一个下标从 0 开始的字符串 s ，它只包含字符 '*' 和 '|' ，其中 '*' 表示一个 盘子 ，'|' 表示一支 蜡烛 。
//
// 同时给你一个下标从 0 开始的二维整数数组 queries ，其中 queries[i] = [lefti, righti] 表示 子字符串 s[lefti...righti] （包含左右端点的字符）。对于每个查询，你需要找到 子字符串中 在 两支蜡烛之间 的盘子的 数目 。如果一个盘子在 子字符串中 左边和右边 都 至少有一支蜡烛，那么这个盘子满足在 两支蜡烛之间 。
//
// 比方说，s = "||**||**|*" ，查询 [3, 8] ，表示的是子字符串 "*||**|" 。子字符串中在两支蜡烛之间的盘子数目为 2 ，子字符串中右边两个盘子在它们左边和右边 都 至少有一支蜡烛。
// 请你返回一个整数数组 answer ，其中 answer[i] 是第 i 个查询的答案。
func platesBetweenCandles(s string, queries [][]int) []int {
	left := make([]int, len(s))
	preSum := make([]int, len(s))
	sum, l := 0, -1
	// 预处理区间内每个位置左侧的第一个蜡烛
	// 计算盘子数量的前缀和 preSum
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			sum++
		} else {
			l = i
		}
		preSum[i] = sum
		left[i] = l
	}

	r := -1
	right := make([]int, len(s))
	// 预处理区间内每个位置右侧的第一个蜡烛
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '|' {
			r = i
		}
		right[i] = r
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		// 找到区间内最左侧和最右侧的蜡烛位置
		x, y := right[q[0]], left[q[1]]
		if x >= 0 && y >= 0 && x < y {
			ans[i] = preSum[y] - preSum[x]
		}
	}
	return ans
}
