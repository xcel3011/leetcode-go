package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	pos := findMaxPos(nums)
	root := &TreeNode{Val: nums[pos]}
	root.Left = constructMaximumBinaryTree(nums[:pos])
	root.Right = constructMaximumBinaryTree(nums[pos+1:])
	return root
}

func findMaxPos(nums []int) int {
	max := 0
	for i := range nums {
		if nums[i] > nums[max] {
			max = i
		}
	}
	return max
}
