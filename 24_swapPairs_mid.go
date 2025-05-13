package leetcode_go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairsK(head *ListNode) *ListNode {
	// 先计算链表长度
	n := 0
	cur := head
	for cur != nil {
		n++
		cur = cur.Next
	}

	// 定义一个虚拟头节点
	dummy := &ListNode{Next: head}
	p0 := dummy
	var pre *ListNode
	cur = p0.Next
	for n >= 2 {
		n -= 2
		// 两个一组反转
		for i := 0; i < 2; i++ {
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

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	p0 := dummy
	p1 := p0.Next
	// [0]->1->2->3->4
	for p1 != nil && p1.Next != nil {
		p2 := p1.Next
		p3 := p2.Next

		p0.Next = p2
		p2.Next = p1
		p1.Next = p3

		p0 = p1
		p1 = p3
	}
	return dummy.Next
}
