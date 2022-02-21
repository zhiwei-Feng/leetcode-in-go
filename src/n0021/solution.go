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
			newTail.Next = p
			p = p.Next
		} else {
			newTail.Next = q
			q = q.Next
		}
		newTail=  newTail.Next
	}

	if p != nil {
		newTail.Next = p
	}
	if q != nil {
		newTail.Next = q
	}
	return newHead.Next
}
