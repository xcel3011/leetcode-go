package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	var (
		traverse func(node *TreeNode)
		vals     []int
	)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		traverse(node.Left)
		traverse(node.Right)
		vals = append(vals, node.Val)
	}

	traverse(root)
	return vals
}
