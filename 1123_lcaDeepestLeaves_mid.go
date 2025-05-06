package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	maxDepth := -1
	var ans *TreeNode
	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, depth int) int {
		if node == nil {
			// 维护全局最大深度
			maxDepth = max(maxDepth, depth)
			return depth
		}

		leftDepth := dfs(node.Left, depth+1)
		rightDepth := dfs(node.Right, depth+1)
		if leftDepth == maxDepth && rightDepth == maxDepth {
			ans = node
		}
		return max(leftDepth, rightDepth)
	}
	dfs(root, 0)
	return ans
}
