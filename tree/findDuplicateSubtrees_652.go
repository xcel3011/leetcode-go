package tree

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var (
	nodes   []*TreeNode
	nodeMap map[string]int
)

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	nodes = make([]*TreeNode, 0)
	nodeMap = make(map[string]int)
	findDuplicateSubtreesHelper(root)
	return nodes
}

func findDuplicateSubtreesHelper(root *TreeNode) string {
	if root == nil {
		return "#"
	}

	left := findDuplicateSubtreesHelper(root.Left)
	right := findDuplicateSubtreesHelper(root.Right)

	subTree := fmt.Sprintf("%s.%s.%d", left, right, root.Val)

	num, ok := nodeMap[subTree]
	if !ok {
		nodeMap[subTree] = 0
	} else {
		num++
		if num == 1 {
			nodes = append(nodes, root)
		}
		nodeMap[subTree] = num
	}

	return subTree
}
