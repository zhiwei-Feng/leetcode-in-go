package n0023

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var (
		finalHead *ListNode
		curHead   *ListNode
		h         MinHeap
	)
	heap.Init(&h)
	for _, head := range lists {
		if head != nil {
			heap.Push(&h, head)
		}
	}
	for h.Len() > 0 {
		small := heap.Pop(&h).(*ListNode)
		if finalHead == nil {
			finalHead = small
			curHead = small
		} else {
			curHead.Next = small
			curHead = curHead.Next
		}
		if small.Next != nil {
			heap.Push(&h, small.Next)
		}
	}

	return finalHead
}

type MinHeap []*ListNode

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func mergeKLists_binary(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return binaryMerge(lists, 0, len(lists)-1)
}

func binaryMerge(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	mid := left + (right-left)/2
	leftList := binaryMerge(lists, left, mid)
	rightList := binaryMerge(lists, mid+1, right)
	return merge(leftList, rightList)
}

func merge(l1, l2 *ListNode) *ListNode {
	p, q := l1, l2
	head := &ListNode{}
	k := head
	for p != nil && q != nil {
		if p.Val <= q.Val {
			k.Next = p
			p = p.Next
		} else {
			k.Next = q
			q = q.Next
		}
		k = k.Next
	}
	for p != nil {
		k.Next = p
		p = p.Next
		k = k.Next
	}
	for q != nil {
		k.Next = q
		q = q.Next
		k = k.Next
	}
	return head.Next
}
