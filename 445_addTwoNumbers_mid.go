package leetcode_go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers445(l1 *ListNode, l2 *ListNode) *ListNode {
	reverse := func(head *ListNode) *ListNode {
		var pre *ListNode
		cur := head
		for cur != nil {
			// 记录当前节点的下一个节点
			nxt := cur.Next
			// 然后将当前节点指向pre
			cur.Next = pre
			// pre和cur节点都前进一位
			pre = cur
			cur = nxt
		}
		return pre
	}
	addTwo := func(l1, l2 *ListNode) *ListNode {
		dummy := &ListNode{}
		cur := dummy
		carry := 0
		for l1 != nil || l2 != nil || carry != 0 {
			if l1 != nil {
				carry += l1.Val
			}
			if l2 != nil {
				carry += l2.Val
			}
			cur.Next = &ListNode{Val: carry % 10}
			carry /= 10
			cur = cur.Next
			if l1 != nil {
				l1 = l1.Next
			}
			if l2 != nil {
				l2 = l2.Next
			}
		}
		return dummy.Next
	}

	return reverse(addTwo(reverse(l1), reverse(l2)))
}
