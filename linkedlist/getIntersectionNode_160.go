package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := 0, 0
	p1, p2 := headA, headB

	for p1 != nil {
		p1 = p1.Next
		lenA++
	}

	for p2 != nil {
		p2 = p2.Next
		lenB++
	}

	if lenA > lenB {
		for i := 0; i < lenA-lenB; i++ {
			headA = headA.Next
		}
	} else {
		for i := 0; i < lenB-lenA; i++ {
			headB = headB.Next
		}
	}

	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}

	return headA
}
