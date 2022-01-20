package offer0025

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newListHead := &ListNode{}
	p, q, r := l1, l2, newListHead
	for p != nil && q != nil {
		if p.Val < q.Val {
			r.Next = p
			p = p.Next
		} else {
			r.Next = q
			q = q.Next
		}
		r = r.Next
	}

	for p != nil {
		r.Next = p
		p = p.Next
		r = r.Next
	}
	for q != nil {
		r.Next = q
		q = q.Next
		r = r.Next
	}
	return newListHead.Next
}
