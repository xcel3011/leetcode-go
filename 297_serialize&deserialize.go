package leetcode_go

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
const (
	SEP  = ","
	NULL = "#"
)

type Codec struct{}

func Constructor() (_ Codec) {
	return
}

// Serializes a tree to a single string.
func (Codec) serialize(root *TreeNode) string {
	var (
		traverse func(node *TreeNode)
		builder  = new(strings.Builder)
	)
	traverse = func(node *TreeNode) {
		if node == nil {
			builder.WriteString(NULL)
			builder.WriteString(SEP)
			return
		}

		builder.WriteString(strconv.Itoa(node.Val))
		builder.WriteString(SEP)

		traverse(node.Left)
		traverse(node.Right)
	}

	traverse(root)
	return builder.String()
}

// Deserializes your encoded data to tree.
func (Codec) deserialize(data string) *TreeNode {
	var (
		traverse func() *TreeNode
		nodes    = strings.Split(data, SEP)
	)

	traverse = func() *TreeNode {
		if len(nodes) == 0 {
			return nil
		}

		if nodes[0] == NULL {
			nodes = nodes[1:]
			return nil
		}

		val, _ := strconv.Atoi(nodes[0])
		root := &TreeNode{Val: val}
		nodes = nodes[1:]

		root.Left = traverse()
		root.Right = traverse()
		return root
	}

	return traverse()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
