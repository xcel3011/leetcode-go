package leetcode_go

import (
	"math"
)

func isBalanced(root *TreeNode) bool {
	var f func(root *TreeNode) int
	f = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := f(root.Left)
		if l == -1 {
			return -1
		}
		r := f(root.Right)
		if r == -1 {
			return -1
		}
		if math.Abs(float64(l-r)) > 1 {
			return -1
		}
		return max(l, r) + 1
	}
	return f(root) != -1
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced250516(root *TreeNode) bool {
	var f func(root *TreeNode) int
	f = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := f(root.Left)
		if l == -1 {
			return -1
		}
		r := f(root.Right)
		if r == -1 || int(math.Abs(float64(l-r))) > 1 {
			return -1
		}
		return max(l, r) + 1
	}
	return f(root) != -1
}
