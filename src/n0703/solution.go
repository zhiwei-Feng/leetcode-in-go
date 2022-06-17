package n0703

import (
	"container/heap"
)

type KthLargest struct {
	H IntMinHeap
	K int
}

func Constructor(k int, nums []int) KthLargest {
	instance := KthLargest{
		H: make(IntMinHeap, 0),
		K: k,
	}
	for _, v := range nums {
		instance.H = append(instance.H, v)
	}
	heap.Init(&instance.H)
	for instance.H.Len() > instance.K {
		heap.Pop(&instance.H)
	}
	return instance
}

func (this *KthLargest) Add(val int) int {
	if this.H.Len() < this.K {
		heap.Push(&this.H, val)
	} else if this.H[0] < val {
		heap.Pop(&this.H)
		heap.Push(&this.H, val)
	}
	return this.H[0]
}

type IntMinHeap []int

func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
