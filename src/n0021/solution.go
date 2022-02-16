package n0021

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	newHead := &ListNode{}
	newTail := newHead
	p, q := list1, list2
	for p != nil && q != nil {
		if p.Val < q.Val {
			pNext := p.Next
			p.Next = nil
			newTail.Next = p
			newTail = newTail.Next
			p = pNext
		} else {
			qNext := q.Next
			q.Next = nil
			newTail.Next = q
			newTail = newTail.Next
			q = qNext
		}
	}

	if p != nil {
		newTail.Next = p
	}
	if q != nil {
		newTail.Next = q
	}
	return newHead.Next
}
