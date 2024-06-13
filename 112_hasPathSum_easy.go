package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	diff := targetSum - root.Val
	if root.Left == nil && root.Right == nil {
		return diff == 0
	}
	if root.Left != nil && hasPathSum(root.Left, diff) {
		return true
	}
	if root.Right != nil && hasPathSum(root.Right, diff) {
		return true
	}
	return false
}
