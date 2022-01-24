package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key == root.Val {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		minMode := root.Right
		for minMode.Left != nil {
			minMode = minMode.Left
		}
		root.Right = deleteNode(root.Right, minMode.Val)
		minMode.Right = root.Right
		minMode.Left = root.Left
		root = minMode
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	}

	return root
}
