package leetcode_go

import (
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) (ans int) {
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, val int) {
		if root == nil {
			return
		}
		if root.Val >= val {
			ans++
			val = root.Val
		}
		dfs(root.Left, val)
		dfs(root.Right, val)
	}
	// 不能从0开始，因为节点的值可以是负数
	dfs(root, math.MinInt)
	return
}
