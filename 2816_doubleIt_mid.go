package leetcode_go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func doubleIt(head *ListNode) *ListNode {
	reverse := func(head *ListNode) *ListNode {
		var pre *ListNode // 初始节点的前一个为空
		cur := head       // 初始节点为头节点
		for cur != nil {
			nxt := cur.Next // 记录当前节点的下一个节点
			cur.Next = pre  // 将当前节点的下一个节点指向前一个节点
			pre = cur       // 将前一个节点指向当前节点
			cur = nxt       // 将当前节点指向下一个节点
		}
		return pre // 返回反转后的头节点
	}

	double := func(head *ListNode) *ListNode {
		dummy := &ListNode{} // 虚拟节点
		cur := dummy         // 初始节点
		carry := 0           // 进位
		for head != nil || carry != 0 {
			if head != nil {
				carry += head.Val * 2
			}
			cur.Next = &ListNode{Val: carry % 10}
			carry /= 10
			cur = cur.Next
			if head != nil {
				head = head.Next
			}
		}
		return dummy.Next
	}
	return reverse(double(reverse(head)))
}
