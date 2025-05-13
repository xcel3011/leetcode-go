package leetcode_go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
	findMid := func(head *ListNode) *ListNode {
		slow, fast := head, head
		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		}
		return slow
	}
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
	mid := findMid(head)
	head2 := reverse(mid)
	ans := 0
	for head2 != nil {
		ans = max(ans, head.Val+head2.Val)
		head = head.Next
		head2 = head2.Next
	}
	return ans
}
