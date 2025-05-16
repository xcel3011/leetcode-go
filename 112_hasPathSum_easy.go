package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	diff := targetSum - root.Val
	if root.Left == nil && root.Right == nil {
		return diff == 0
	}
	if root.Left != nil && hasPathSum(root.Left, diff) {
		return true
	}
	if root.Right != nil && hasPathSum(root.Right, diff) {
		return true
	}
	return false
}

func hasPathSum250516(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	// 如果遇到叶子节点，判断是否满足条件
	if root.Left == nil && root.Right == nil {
		return targetSum == 0
	}
	return hasPathSum250516(root.Left, targetSum) || hasPathSum250516(root.Right, targetSum)
}
