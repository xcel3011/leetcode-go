package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var vals []int

	vals = append(vals, inorderTraversal(root.Left)...)
	vals = append(vals, root.Val)
	vals = append(vals, inorderTraversal(root.Right)...)

	return vals
}
