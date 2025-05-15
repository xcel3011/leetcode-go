package leetcode_go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNodes(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	node := removeNodes(head.Next)
	if node.Val > head.Val {
		return node
	}
	head.Next = node
	return head
}
