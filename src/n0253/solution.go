package n0253

import (
	"container/heap"
	"sort"
)

func minMeetingRooms(intervals [][]int) int {
    start_time := make([]int, len(intervals))
    end_time := make([]int, len(intervals))
    for i, v:= range intervals{
        start_time[i]=v[0]
        end_time[i]=v[1]
    }
    sort.Ints(start_time)
    sort.Ints(end_time)

    usedRoom := 0
    s_ptr, e_ptr := 0, 0
    for s_ptr<len(intervals) {
        if start_time[s_ptr]>=end_time[e_ptr]{
            usedRoom--
            e_ptr++
        }
        usedRoom++
        s_ptr++
    }
    return usedRoom
}

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func minMeetingRooms_(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	var minH MinHeap
	heap.Init(&minH)
	sort.SliceStable(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		} else {
			return intervals[i][1] < intervals[j][1]
		}
	})
	maxExits := 0
	for i := range intervals {
		if minH.Len() == 0 || minH[0] > intervals[i][0] {
			heap.Push(&minH, intervals[i][1])
		} else {
			heap.Pop(&minH)
			heap.Push(&minH, intervals[i][1])
		}
		maxExits = max(maxExits, minH.Len())
	}
	return maxExits
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
