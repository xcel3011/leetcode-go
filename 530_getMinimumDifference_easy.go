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
func getMinimumDifference(root *TreeNode) int {
	pre := math.MinInt / 2
	ans := math.MaxInt
	var f func(node *TreeNode)
	f = func(node *TreeNode) {
		if node == nil {
			return
		}
		f(node.Left)
		ans = min(ans, node.Val-pre)
		pre = node.Val
		f(node.Right)
	}
	f(root)
	return ans
}
