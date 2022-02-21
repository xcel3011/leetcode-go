package linkedlist

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 在两条链表上的指针
	p1, p2 := l1, l2
	// 虚拟头结点（构建新链表时的常用技巧）
	dummy := new(ListNode)
	// 指针 p 负责构建新链表
	p := dummy
	// 记录进位
	carry := 0
	// 开始执行加法，两条链表走完且没有进位时才能结束循环
	for p1 != nil || p2 != nil || carry != 0 {
		// 先加上上次的进位
		val := carry

		if p1 != nil {
			val += p1.Val
			p1 = p1.Next
		}

		if p2 != nil {
			val += p2.Val
			p2 = p2.Next
		}

		carry = val / 10
		val = val % 10
		p.Next = &ListNode{Val: val}
		p = p.Next
	}
	return dummy.Next
}
