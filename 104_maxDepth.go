package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var (
	depth int
	res   int
)

func maxDepth(root *TreeNode) int {
	depth, res = 0, 0
	traverse(root)
	return res
}

func traverse(root *TreeNode) {
	if root == nil {
		res = max(res, depth)
		return
	}

	depth++
	traverse(root.Left)
	traverse(root.Right)
	depth--
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
