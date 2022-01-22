package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree1(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	index := 0
	for i, val := range inorder {
		if val == postorder[len(postorder)-1] {
			index = i
			break
		}
	}

	root := &TreeNode{Val: inorder[index]}
	root.Left = buildTree1(inorder[:index], postorder[:index])
	root.Right = buildTree1(inorder[index+1:], postorder[index:len(postorder)-1])
	return root
}
