package leetcode_go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	slow, fast := dummy, dummy

	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}

func removeNthFromEnd050515(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}

	// 删除倒数N个节点要找到倒数N+1个节点
	// 右指针先走N步
	right := dummy
	for i := 0; i < n; i++ {
		right = right.Next
	}
	// 保证左右指针相差N+1步
	left := dummy
	// 因为有dummy，当右指针走到最后一个节点的时候，左指针刚好在倒数N+1个节点
	for right.Next != nil {
		right = right.Next
		left = left.Next
	}

	// 删除倒数N个节点
	left.Next = left.Next.Next
	return dummy.Next
}
