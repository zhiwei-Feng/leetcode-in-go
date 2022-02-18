package n0019

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var prev *ListNode
	p, q := head, head
	for n != 0 && q != nil {
		q = q.Next
		n--
	}
	for q != nil {
		prev = p
		p = p.Next
		q = q.Next
	}

	if prev == nil {
		return head.Next
	}
	prev.Next = p.Next
	return head
}
