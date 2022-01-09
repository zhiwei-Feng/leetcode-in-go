package offer0024

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	p := head
	reverseHead := &ListNode{}
	for p != nil {
		nextNode := p.Next
		p.Next = reverseHead.Next
		reverseHead.Next = p
		p = nextNode
	}
	return reverseHead.Next
}
