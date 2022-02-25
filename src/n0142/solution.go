package n0142

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	fast = dummy
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
