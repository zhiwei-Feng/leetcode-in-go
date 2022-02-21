package n0148

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
    dummy := &ListNode{Next:head}
	mid, p := dummy, dummy
	for p != nil && p.Next != nil {
		p = p.Next.Next
		mid = mid.Next
	}
	rHead := mid.Next
	mid.Next = nil
	leftList, rightList := sortList(head), sortList(rHead)
	return mergeTwoLists(leftList, rightList)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	newHead := &ListNode{}
	newTail := newHead
	p, q := list1, list2
	for p != nil && q != nil {
		if p.Val < q.Val {
			newTail.Next = p
			p = p.Next
		} else {
			newTail.Next = q
			q = q.Next
		}
		newTail=  newTail.Next
	}

	if p != nil {
		newTail.Next = p
	}
	if q != nil {
		newTail.Next = q
	}
	return newHead.Next
}

func sortList_(head *ListNode) *ListNode {
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
			for prev.Next.Val <= curNode.Val {
				prev = prev.Next
			}
			lastOrderedNode.Next = curNode.Next
			curNode.Next = prev.Next
			prev.Next = curNode
		}
	}

	return H.Next
}
