package n0215

import (
	"container/heap"
	"math/rand"
)

type minHeap []int

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// 用一个k容量的小顶堆存储最大的k个数，此时堆顶为目标
// O(nlogn)
func findKthLargest_(nums []int, k int) int {
	var mh minHeap
	heap.Init(&mh)

	for _, v := range nums {
		if mh.Len() < k {
			heap.Push(&mh, v)
			continue
		}

		if mh[0] < v {
			heap.Pop(&mh)
			heap.Push(&mh, v)
		}
	}

	return mh[0]
}

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	if k > n {
		return -1
	}
	low, high := 0, n-1
	for {
		index := partion(nums, low, high)
		if index == k-1 {
			return nums[index]
		} else if index < k-1 {
			low = index + 1
		} else {
			high = index - 1
		}
	}
}

func partion(nums []int, low, high int) int {
	if low == high {
		return low
	}

	randIndex := rand.Intn(high-low+1) + low
	nums[low], nums[randIndex] = nums[randIndex], nums[low]
	pivot := nums[low]
	j := low
	for i := low + 1; i <= high; i++ {
		if nums[i] > pivot {
			j++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[low], nums[j] = nums[j], nums[low]
	return j
}
