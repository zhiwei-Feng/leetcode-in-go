package n0373

import "container/heap"

type minHeap [][]int // []int{i, j, sum}

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i][2] < h[j][2] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	m, n := len(nums1), len(nums2)
	minH := minHeap{}
	for i := 0; i < m && i < k; i++ {
		minH = append(minH, []int{i, 0, nums1[i] + nums2[0]})
	}
	heap.Init(&minH)
	ans := [][]int{}
	for len(minH) > 0 && len(ans) < k {
		top := heap.Pop(&minH).([]int)
		i, j := top[0], top[1]
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j+1 < n {
			heap.Push(&minH, []int{i, j + 1, nums1[i]+nums2[j+1]})
		}
	}

	return ans
}
