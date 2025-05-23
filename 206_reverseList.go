package leetcode_go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// /**
// * 以链表1->2->3->4->5举例
// * @param head
// * @return
// */
// public ListNode reverseList(ListNode head) {
// if (head == null || head.next == null) {
// /*
//
//	直到当前节点的下一个节点为空时返回当前节点
//	由于5没有下一个节点了，所以此处返回节点5
//
// */
// return head;
// }
// //递归传入下一个节点，目的是为了到达最后一个节点
// ListNode newHead = reverseList(head.next);
// /*
//
//	第一轮出栈，head为5，head.next为空，返回5
//	第二轮出栈，head为4，head.next为5，执行head.next.next=head也就是5.next=4，
//	          把当前节点的子节点的子节点指向当前节点
//	          此时链表为1->2->3->4<->5，由于4与5互相指向，所以此处要断开4.next=null
//	          此时链表为1->2->3->4<-5
//	          返回节点5
//	第三轮出栈，head为3，head.next为4，执行head.next.next=head也就是4.next=3，
//	          此时链表为1->2->3<->4<-5，由于3与4互相指向，所以此处要断开3.next=null
//	          此时链表为1->2->3<-4<-5
//	          返回节点5
//	第四轮出栈，head为2，head.next为3，执行head.next.next=head也就是3.next=2，
//	          此时链表为1->2<->3<-4<-5，由于2与3互相指向，所以此处要断开2.next=null
//	          此时链表为1->2<-3<-4<-5
//	          返回节点5
//	第五轮出栈，head为1，head.next为2，执行head.next.next=head也就是2.next=1，
//	          此时链表为1<->2<-3<-4<-5，由于1与2互相指向，所以此处要断开1.next=null
//	          此时链表为1<-2<-3<-4<-5
//	          返回节点5
//	出栈完成，最终头节点5->4->3-2->1
//
// */
// head.next.next = head;
// head.next = null;
// return newHead;
// }
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return last
}

func reverseListIterate(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func reverseList250513(head *ListNode) *ListNode {
	// 当前head的前一个节点为空
	var pre *ListNode
	cur := head

	for cur != nil {
		// 记录当前节点的下一个节点
		nxt := cur.Next
		// 当前节点的下一个节点指向当前节点的前一个节点
		cur.Next = pre
		// 当前节点的前一个节点指向当前节点
		pre = cur
		// 当前节点指向当前节点的下一个节点
		cur = nxt
	}
	return pre
}
