package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	ans := 0
	var dfs func(root *TreeNode, x int)
	dfs = func(root *TreeNode, x int) {
		if root == nil {
			return
		}
		x = x*10 + root.Val
		// 叶子节点
		if root.Left == nil && root.Right == nil {
			ans += x
			return
		}
		dfs(root.Left, x)
		dfs(root.Right, x)
	}
	dfs(root, 0)
	return ans
}
