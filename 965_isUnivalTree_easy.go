package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
	var dfs func(root *TreeNode, val int) bool
	dfs = func(root *TreeNode, val int) bool {
		if root == nil {
			return true
		}
		if root.Val != val {
			return false
		}
		return dfs(root.Left, val) && dfs(root.Right, val)
	}
	return dfs(root, root.Val)
}
