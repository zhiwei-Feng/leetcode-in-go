package n0206

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	newHead := &ListNode{}
	p := head
	for p != nil {
		pNext := p.Next
		p.Next = newHead.Next
		newHead.Next = p
		p = pNext
	}
	return newHead.Next
}
