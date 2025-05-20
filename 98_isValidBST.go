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
func isValidBST(root *TreeNode) bool {
	var traverse func(node, min, max *TreeNode) bool
	traverse = func(node, min, max *TreeNode) bool {
		if node == nil {
			return true
		}
		if min != nil && min.Val >= node.Val {
			return false
		}
		if max != nil && max.Val <= node.Val {
			return false
		}

		return traverse(node.Left, min, node) && traverse(node.Right, node, max)
	}
	return traverse(root, nil, nil)
}

// 前序遍历
func isValidBSTPreOrder250520(root *TreeNode) bool {
	var f func(node *TreeNode, mx, mn int) bool
	f = func(node *TreeNode, mx, mn int) bool {
		if node == nil {
			return true
		}

		return node.Val > mn && node.Val < mx && f(node.Left, node.Val, mn) && f(node.Right, mx, node.Val)
	}
	return f(root, math.MaxInt, math.MinInt)
}

// 中序遍历
func isValidBSTInOrder250520(root *TreeNode) bool {
	pre := math.MinInt
	var f func(node *TreeNode) bool
	f = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if !f(node.Left) {
			return false
		}
		if node.Val <= pre {
			return false
		}
		pre = node.Val
		return f(node.Right)
	}
	return f(root)
}

// 后序遍历
func isValidBSTPostOrder250520(root *TreeNode) bool {
	var f func(node *TreeNode) (mx, mn int)
	f = func(node *TreeNode) (mx, mn int) {
		if node == nil {
			return math.MinInt, math.MaxInt
		}
		lmx, lmn := f(node.Left)
		rmx, rmn := f(node.Right)
		if node.Val <= lmx || node.Val >= rmn {
			return math.MaxInt, math.MinInt
		}
		return max(rmx, node.Val), min(lmn, node.Val)
	}
	mx, _ := f(root)
	return mx != math.MaxInt
}
