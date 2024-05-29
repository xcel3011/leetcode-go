package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	var helper func(lo, hi int) []*TreeNode
	helper = func(lo, hi int) []*TreeNode {
		var res []*TreeNode
		if lo > hi {
			return append(res, nil)
		}

		for i := lo; i <= hi; i++ {
			left := helper(lo, i-1)
			right := helper(i+1, hi)
			for _, leftNode := range left {
				for _, rightNode := range right {
					root := &TreeNode{Val: i}
					root.Left = leftNode
					root.Right = rightNode
					res = append(res, root)
				}
			}
		}
		return res
	}

	return helper(1, n)
}
