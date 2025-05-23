package leetcode_go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome234(head *ListNode) bool {
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

func isPalindrome250513(head *ListNode) bool {
	reverse := func(head *ListNode) *ListNode {
		var pre *ListNode
		cur := head
		for cur != nil {
			nxt := cur.Next
			cur.Next = pre
			pre = cur
			cur = nxt
		}
		return pre
	}
	// 找到中间节点
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
		left = left.Next
		right = right.Next
	}
	return true
}
