package offer0018

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	p := head
	for p!=nil && p.Next != nil {
		deleteNext := p.Next.Next
		if p.Next.Val == val {
			// delete p.Next
			p.Next = p.Next.Next
			p = deleteNext
		} else {
			p = p.Next
		}
	}
	return head
}
