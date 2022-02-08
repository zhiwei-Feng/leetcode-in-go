package offer0059_1

import "container/heap"

func maxSlidingWindow(nums []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	ans := make([]int, 0, len(nums)/k)
	deque := make([]int, 0, k)

	i := 0
	for j := 0; j < len(nums); j++ {
		for len(deque) > 0 && deque[len(deque)-1] < nums[j] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, nums[j])
		if j < k-1 {
			continue
		}

		ans = append(ans, deque[0])
		if deque[0] == nums[i] {
			deque = deque[1:]
		}
		i++
	}
	return ans
}

type pair struct {
	Val int
	Ind int
}

type MaxHeap []pair

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i].Val > h[j].Val }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(pair)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maxSlidingWindow_slidingwindow(nums []int, k int) []int {
	winStart := 0
	var winHeap MaxHeap
	heap.Init(&winHeap)
	ans := []int{}
	for winEnd := 0; winEnd < len(nums); winEnd++ {
		heap.Push(&winHeap, pair{nums[winEnd], winEnd})
		if winEnd < k-1 {
			continue
		}
		for winHeap[0].Ind < winStart {
			heap.Pop(&winHeap)
		}
		ans = append(ans, winHeap[0].Val)
		winStart++
	}

	return ans
}
