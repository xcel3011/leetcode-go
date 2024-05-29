package leetcode_go

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxSumBST(root *TreeNode) int {
	var (
		traverse func(node *TreeNode) [4]int
		max      func(a, b int) int
		min      func(a, b int) int
		maxSum   int
	)

	min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 0 标记是否为二叉搜索树
	// 1 标记最小值
	// 2 标记最大值
	// 3 标记和
	traverse = func(node *TreeNode) [4]int {
		if node == nil {
			return [4]int{1, math.MaxInt, math.MinInt, 0}
		}

		// 递归遍历左右子树
		left := traverse(node.Left)
		right := traverse(node.Right)

		// 后续遍历
		// 左右子树要都是二叉搜索树
		// 根节点大于左子树的最大值、小于右子树的最小值
		var res [4]int
		if left[0] == 1 && right[0] == 1 && node.Val > left[2] && node.Val < right[1] {
			res[0] = 1
			res[1] = min(left[1], node.Val)
			res[2] = max(right[2], node.Val)
			res[3] = left[3] + right[3] + node.Val
			maxSum = max(maxSum, res[3])
		} else {
			res[0] = 0
		}
		return res
	}

	traverse(root)
	return maxSum
}
