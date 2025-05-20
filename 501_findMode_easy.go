package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) (ans []int) {
	var base, cnt, mxcnt int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if base == root.Val {
			cnt++
		} else {
			base, cnt = root.Val, 1
		}

		if cnt == mxcnt {
			ans = append(ans, base)
		} else if cnt > mxcnt {
			mxcnt = cnt
			ans = []int{base}
		}

		dfs(root.Right)
	}
	dfs(root)
	return ans
}
