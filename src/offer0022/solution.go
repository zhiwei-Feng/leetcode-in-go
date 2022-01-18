package offer0022

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	p := head
	q := head
	for k > 0 && q != nil {
		q = q.Next
		k--
	}
	if k != 0 {
		return nil
	}

	for q != nil {
		p = p.Next
		q = q.Next
	}
	return p
}
