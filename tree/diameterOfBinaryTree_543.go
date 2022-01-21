package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var diameter int

func diameterOfBinaryTree(root *TreeNode) int {
	diameter = 0
	diameterHelper(root)
	return diameter
}

func diameterHelper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDiameter := diameterHelper(root.Left)
	rightDiameter := diameterHelper(root.Right)
	myDiameter := leftDiameter + rightDiameter
	diameter = max(myDiameter, diameter)

	return 1 + max(leftDiameter, rightDiameter)
}
