package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	dummy := new(ListNode)
	p := dummy
	for head != nil {
		if head.Val != val {
			p.Next = &ListNode{Val: head.Val}
			p = p.Next
		}
		head = head.Next
	}

	return dummy.Next
}
