package leetcode_go

import (
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepthOld(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var (
		queue []*TreeNode
		depth = 1
	)

	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left == nil && node.Right == nil {
				return depth
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
	}
	return depth
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	depth := math.MaxInt
	if root.Left != nil {
		depth = min(minDepth(root.Left), depth)
	}
	if root.Right != nil {
		depth = min(minDepth(root.Right), depth)
	}

	return depth + 1
}
