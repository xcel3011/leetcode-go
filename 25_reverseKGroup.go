package leetcode_go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}

	res := reverse(a, b)
	a.Next = reverseKGroup(b, k)

	return res
}

func reverse(a, b *ListNode) *ListNode {
	var prev *ListNode
	curr := a

	for curr != b {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
