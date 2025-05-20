package leetcode_go

import (
	"slices"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	index := 0
	for i := range inorder {
		if preorder[0] == inorder[i] {
			index = i
			break
		}
	}

	root := new(TreeNode)
	root.Val = preorder[0]
	root.Left = buildTree(preorder[1:index+1], inorder[:index])
	root.Right = buildTree(preorder[index+1:], inorder[index+1:])
	return root
}

func buildTree250520(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	idx := slices.Index(inorder, preorder[0])
	root := &TreeNode{Val: preorder[0]}
	root.Left = buildTree250520(preorder[1:idx+1], inorder[:idx])
	root.Right = buildTree250520(preorder[idx+1:], inorder[idx+1:])
	return root
}
