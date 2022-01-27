package offer0041

import "container/heap"

type MaxHeap []int
type MinHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{}   { old := *h; x := old[len(old)-1]; *h = old[:len(old)-1]; return x }

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{}   { old := *h; x := old[len(old)-1]; *h = old[:len(old)-1]; return x }

type MedianFinder struct {
	AHeap MaxHeap // 存较小的一半
	BHeap MinHeap // 存较大的一半
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	var ah MaxHeap
	var bh MinHeap
	heap.Init(&ah)
	heap.Init(&bh)
	return MedianFinder{AHeap: ah, BHeap: bh}
}

func (this *MedianFinder) AddNum(num int) {
	m, n := len(this.AHeap), len(this.BHeap)
	if m != n {
		// 插入A，弹顶入B
		heap.Push(&this.AHeap, num)
		x := heap.Pop(&this.AHeap)
		heap.Push(&this.BHeap, x)
	}else {
		// 插入B，弹顶入A
		heap.Push(&this.BHeap, num)
		x := heap.Pop(&this.BHeap)
		heap.Push(&this.AHeap, x)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	m, n := len(this.AHeap), len(this.BHeap)
	if m == n {
		return (float64(this.AHeap[0]) + float64(this.BHeap[0])) / 2
	} else {
		return float64(this.AHeap[0])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
