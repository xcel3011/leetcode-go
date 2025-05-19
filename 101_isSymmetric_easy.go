package leetcode_go

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var f func(*TreeNode, *TreeNode) bool
	f = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		return f(left.Left, right.Right) && f(left.Right, right.Left)
	}
	return f(root.Left, root.Right)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric250516(root *TreeNode) bool {
	var f func(p, q *TreeNode) bool
	f = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}

		return p.Val == q.Val && f(p.Left, q.Right) && f(p.Right, q.Left)
	}
	return f(root.Left, root.Right)
}
