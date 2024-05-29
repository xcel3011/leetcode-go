package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 分解子树的最大高度
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth1(root.Left)
	rightDepth := maxDepth1(root.Right)

	return max(leftDepth, rightDepth) + 1
}
