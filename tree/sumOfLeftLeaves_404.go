package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var count int
	if root.Left != nil &&
		root.Left.Left == nil &&
		root.Left.Right == nil {
		count += root.Left.Val
	}

	count += sumOfLeftLeaves(root.Left)
	count += sumOfLeftLeaves(root.Right)

	return count
}
