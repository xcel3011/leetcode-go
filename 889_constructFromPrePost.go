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
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	index := -1
	for i, val := range postorder {
		if val == preorder[1] {
			index = i
			break
		}
	}

	root := &TreeNode{Val: preorder[0]}
	root.Left = constructFromPrePost(preorder[1:index+2], postorder[:index+1])
	root.Right = constructFromPrePost(preorder[index+2:], postorder[index+1:len(postorder)-1])
	return root
}

func constructFromPrePost250520(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	// 前序是根左右，后序是左右根
	index := slices.Index(postorder, preorder[1])
	root := &TreeNode{Val: preorder[0]}
	root.Left = constructFromPrePost250520(preorder[1:index+2], postorder[:index+1])
	root.Right = constructFromPrePost250520(preorder[index+2:], postorder[index+1:len(postorder)-1])
	return root
}
