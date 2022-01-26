package offer0040

import "container/heap"

type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{}   { old := *h; x := old[len(old)-1]; *h = old[:len(old)-1]; return x }

func getLeastNumbers(arr []int, k int) []int {
	if k == 0 || len(arr) == 0 {
		return []int{}
	}
	var hp MaxHeap
	heap.Init(&hp)
	for _, v := range arr {
		if hp.Len() < k {
			heap.Push(&hp, v)
		} else {
			if hp[0] > v {
				heap.Pop(&hp)
				heap.Push(&hp, v)
			}
		}
	}
	return hp
}
