package leetcode_go

import (
	"slices"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestNodes(root *TreeNode, queries []int) [][]int {
	// 中序遍历
	var nums []int
	var f func(root *TreeNode)
	f = func(root *TreeNode) {
		if root == nil {
			return
		}
		f(root.Left)
		nums = append(nums, root.Val)
		f(root.Right)
	}
	f(root)

	// 二分查找
	ans := make([][]int, len(queries))
	for i, q := range queries {
		mx, mn := -1, -1
		j, ok := slices.BinarySearch(nums, q)
		if j < len(nums) {
			mx = nums[j]
		}
		if !ok { // a[j]>q, a[j-1]<q
			j--
		}
		if j >= 0 {
			mn = nums[j]
		}
		ans[i] = []int{mn, mx}
	}
	return ans
}
