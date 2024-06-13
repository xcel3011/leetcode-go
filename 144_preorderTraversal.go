package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var vals []int

func preorderTraversal(root *TreeNode) []int {
	vals = vals[:0]
	preorderTraversalHelper(root)
	return vals
}

func preorderTraversalHelper(root *TreeNode) {
	if root == nil {
		return
	}

	vals = append(vals, root.Val)
	preorderTraversalHelper(root.Left)
	preorderTraversalHelper(root.Right)
}

func preorderTraversal20240612(root *TreeNode) []int {
	var (
		vals []int
		f    func(node *TreeNode)
	)
	f = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		f(node.Left)
		f(node.Right)
	}
	f(root)
	return vals
}
