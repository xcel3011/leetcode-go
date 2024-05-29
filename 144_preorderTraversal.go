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
