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

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var res []*TreeNode
	if root == nil {
		return res
	}

	subTreeMap := make(map[string]int)
	var dfs func(*TreeNode) string
	dfs = func(node *TreeNode) string {
		if node == nil {
			return "#"
		}
		subTree := fmt.Sprintf("%v.%v.%v", dfs(node.Left), dfs(node.Right), node.Val)
		subTreeMap[subTree]++
		if subTreeMap[subTree] == 2 {
			res = append(res, node)
		}
		return subTree
	}
	dfs(root)
	return res
}
