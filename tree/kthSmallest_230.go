package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	index, val := 0, 0

	var traverse func(*TreeNode, int)
	traverse = func(node *TreeNode, k int) {
		if node == nil {
			return
		}

		traverse(node.Left, k)
		index++
		if index == k {
			val = node.Val
		}
		traverse(node.Right, k)
	}
	traverse(root, k)
	return val
}
