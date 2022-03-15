package bytedance0003

import "container/heap"

// 有N个窗口，给定每个窗口处理一个人所需的时间。求M个人都处理完的最短时间
// N=3, M=2, window_time=[2, 2, 5]

func solve(N, M int, winTime []int) int {
	var hp MinHeap
	heap.Init(&hp)
	for i := 0; i < N; i++ {
		heap.Push(&hp, Window{winTime[i], 0})
	}
	// 处理M>N的情况
	for i := 0; i < M; i++ {
		enableWindow := heap.Pop(&hp).(Window)
		enableWindow.EnableTime += enableWindow.HandleTime
		heap.Push(&hp, enableWindow)
	}
	// 求最大的enableTime
	lastTime := 0
	for i := 0; i < len(hp); i++ {
		if hp[i].EnableTime > lastTime {
			lastTime = hp[i].EnableTime
		}
	}
	return lastTime
}

type Window struct {
	HandleTime int
	EnableTime int
}

type MinHeap []Window

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	return h[i].EnableTime+h[i].HandleTime < h[j].EnableTime+h[j].HandleTime
}
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Window)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
