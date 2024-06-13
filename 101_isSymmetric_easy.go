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
