package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	tempNode := root.Left
	root.Left = root.Right
	root.Right = tempNode

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

func invertTree250516(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := root.Left
	root.Left = root.Right
	root.Right = left
	invertTree250516(root.Left)
	invertTree250516(root.Right)
	return root
}
