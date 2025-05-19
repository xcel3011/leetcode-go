package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	if root == nil {
		return nil
	}
	limit -= root.Val

	// 当root为叶子节点
	if root.Left == nil && root.Right == nil {
		if limit > 0 {
			return nil
		}
		return root
	}
	root.Left = sufficientSubset(root.Left, limit)
	root.Right = sufficientSubset(root.Right, limit)
	// 如果儿子都被删除，就删 root
	if root.Left == nil && root.Right == nil {
		return nil
	}
	return root
}
