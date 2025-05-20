package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > val {
		root = searchBST(root.Left, val)
	} else if root.Val < val {
		root = searchBST(root.Right, val)
	}
	return root
}

func searchBST250520(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > val {
		return searchBST250520(root.Left, val)
	} else if root.Val < val {
		return searchBST250520(root.Right, val)
	}
	return root
}
