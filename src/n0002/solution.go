package n0002

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (root *ListNode) {
	carry := 0
	p, q := l1, l2
	var r *ListNode
	for p != nil && q != nil {
		val := (p.Val + q.Val + carry) % 10
		carry = (p.Val + q.Val + carry) / 10
		if root == nil {
			root = &ListNode{Val: val}
			r = root
		} else {
			r.Next = &ListNode{Val: val}
			r = r.Next
		}
		p = p.Next
		q = q.Next
	}

	for p != nil {
		val := (carry + p.Val) % 10
		carry = (carry + p.Val) / 10
		r.Next = &ListNode{Val: val}
		r = r.Next
		p = p.Next
	}
	for q != nil {
		val := (carry + q.Val) % 10
		carry = (carry + q.Val) / 10
		r.Next = &ListNode{Val: val}
		r = r.Next
		q = q.Next
	}
	if carry != 0 {
		r.Next = &ListNode{Val: carry}
	}
	return root
}
