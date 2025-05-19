package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) int {
	ans := 0
	var dfs func(root *TreeNode) (int, int)
	dfs = func(root *TreeNode) (int, int) {
		if root == nil {
			return -1, -1
		} else if root.Left == nil && root.Right == nil {
			return 0, 0
		}

		l, _ := dfs(root.Left)
		_, r := dfs(root.Right)
		l, r = l+1, r+1
		ans = max(ans, l, r)
		return l, r
	}
	dfs(root)
	return ans
}
