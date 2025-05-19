package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxAncestorDiff(root *TreeNode) int {
	ans := 0
	var dfs func(root *TreeNode, mx, mn int)
	dfs = func(root *TreeNode, mx, mn int) {
		if root == nil {
			return
		}
		mx = max(mx, root.Val)
		mn = min(mn, root.Val)
		ans = max(ans, mx-root.Val, root.Val-mn)
		dfs(root.Left, mx, mn)
		dfs(root.Right, mx, mn)
	}
	dfs(root, root.Val, root.Val)
	return ans
}
