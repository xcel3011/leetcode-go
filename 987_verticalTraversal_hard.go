package leetcode_go

import (
	"slices"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func verticalTraversal(root *TreeNode) [][]int {
	type pair struct{ row, val int }
	groups := make(map[int][]pair)
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, row, col int) {
		if node == nil {
			return
		}
		// 记录当前列的值
		groups[col] = append(groups[col], pair{row: row, val: node.Val})
		dfs(node.Left, row+1, col-1)
		dfs(node.Right, row+1, col+1)
	}
	dfs(root, 0, 0)
	// 按照列排序
	cols := make([]int, 0, len(groups))
	for col := range groups {
		cols = append(cols, col)
	}
	slices.Sort(cols)
	ans := make([][]int, len(cols))
	for i, col := range cols {
		rows := groups[col]
		slices.SortFunc(rows, func(a, b pair) int {
			if a.row != b.row {
				return a.row - b.row
			}
			return a.val - b.val
		})
		ans[i] = make([]int, len(rows))
		for j, row := range rows {
			ans[i][j] = row.val
		}
	}
	return ans
}
