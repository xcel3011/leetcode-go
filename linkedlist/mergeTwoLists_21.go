package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	ln := new(ListNode)
	dummy := ln
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			ln.Next = l2
			l2 = l2.Next
		} else {
			ln.Next = l1
			l1 = l1.Next
		}
		ln = ln.Next
	}
	if l1 != nil {
		ln.Next = l1
	}
	if l2 != nil {
		ln.Next = l2
	}
	return dummy.Next
}
