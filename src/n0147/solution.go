package n0147

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	H := &ListNode{Next: head}
	lastOrderedNode := H.Next
	for lastOrderedNode.Next != nil {
		curNode := lastOrderedNode.Next
		if curNode.Val >= lastOrderedNode.Val {
			// 插入最后
			lastOrderedNode = curNode
			continue
		} else {
			prev := H
            for prev.Next.Val<=curNode.Val {
                prev = prev.Next
            }
			lastOrderedNode.Next = curNode.Next
			curNode.Next = prev.Next
			prev.Next = curNode
		}
	}
	
	return H.Next
}
