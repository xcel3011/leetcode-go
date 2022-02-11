package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	reverse := func(node *ListNode) *ListNode {
		var prev *ListNode
		curr := node

		for curr != nil {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}

		return prev
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	left := head
	right := reverse(slow)

	for right != nil {
		if left.Val != right.Val {
			return false
		}
		right = right.Next
		left = left.Next
	}

	return true
}
