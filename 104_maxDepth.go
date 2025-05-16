package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	var f func(root *TreeNode)
	ans, depth := 0, 0
	f = func(root *TreeNode) {
		if root == nil {
			ans = max(ans, depth)
			return
		}
		depth++
		f(root.Left)
		f(root.Right)
		depth--
	}
	f(root)
	return ans
}

func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth1(root.Left)
	rightDepth := maxDepth1(root.Right)

	return max(leftDepth, rightDepth) + 1
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth250516(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth250516(root.Left)
	right := maxDepth250516(root.Right)
	return max(right, left) + 1
}
