package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	var traverse func(node, min, max *TreeNode) bool
	traverse = func(node, min, max *TreeNode) bool {
		if node == nil {
			return true
		}
		if min != nil && min.Val >= node.Val {
			return false
		}
		if max != nil && max.Val <= node.Val {
			return false
		}

		return traverse(node.Left, min, node) && traverse(node.Right, node, max)
	}
	return traverse(root, nil, nil)
}
