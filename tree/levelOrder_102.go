package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var (
		queue   []*TreeNode
		node    *TreeNode
		results [][]int
	)

	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		var result []int
		for i := 0; i < size; i++ {
			node, queue = queue[0], queue[1:]
			result = append(result, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		results = append(results, result)
	}

	return results
}
