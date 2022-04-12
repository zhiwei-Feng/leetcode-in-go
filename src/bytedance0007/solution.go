package bytedance0007

// 给定一个奇数位升序，偶数位降序的链表，将其重新排序。
// 输入: 1->8->3->6->5->4->7->2->NULL
// 输出: 1->2->3->4->5->6->7->8->NULL

type ListNode struct {
	Val  int
	Next *ListNode
}

func solve(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	oddDummy := &ListNode{}
	evenDummy := &ListNode{}
	oddLast := oddDummy
	evenLast := evenDummy

	// 1. 拆分
	p := head
	pos := 1
	for p != nil {
		if pos%2 == 1 {
			nextp := p.Next
			p.Next = nil
			oddLast.Next = p
			oddLast = oddLast.Next
			p = nextp
		} else {
			nextp := p.Next
			p.Next = nil
			evenLast.Next = p
			evenLast = evenLast.Next
			p = nextp
		}
		pos++
	}
	// 2. 逆转evenLast
	evenDummy.Next = reverse(evenDummy.Next)
	// 3. Merge Two List
	return merge(oddDummy.Next, evenDummy.Next)
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	p := head
	for p != nil {
		nextp := p.Next
		p.Next = prev
		prev = p
		p = nextp
	}
	return prev
}

func merge(a, b *ListNode) *ListNode {
	dummy := &ListNode{}
	last := dummy
	for a != nil && b != nil {
		if a.Val <= b.Val {
			nexta := a.Next
			last.Next = a
			last = last.Next
			a = nexta
		} else {
			nextb := b.Next
			last.Next = b
			last = last.Next
			b = nextb
		}
	}
	for a != nil {
		nexta := a.Next
		last.Next = a
		last = last.Next
		a = nexta
	}
	for b != nil {
		nextb := b.Next
		last.Next = b
		last = last.Next
		b = nextb
	}
	return dummy.Next
}
