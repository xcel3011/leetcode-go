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

func reverseKGroup250513(head *ListNode, k int) *ListNode {
	// 先计算链表长度
	n := 0
	cur := head
	for cur != nil {
		n++
		cur = cur.Next
	}

	dummy := &ListNode{Next: head}
	p0 := dummy
	var pre *ListNode
	cur = p0.Next
	for n >= k {
		n -= k
		// 反转k个节点
		for i := 0; i < k; i++ {
			nxt := cur.Next
			cur.Next = pre
			pre = cur
			cur = nxt
		}

		nxt := p0.Next
		p0.Next.Next = cur
		p0.Next = pre
		p0 = nxt
	}
	return dummy.Next
}
