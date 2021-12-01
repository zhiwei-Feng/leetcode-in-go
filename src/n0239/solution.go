package n0239

import ()


// monotonous queue method: time=O(n), space=O(k)
// monoQueue store element index, and the corresponding value is monotonous(left side is bigger)
func maxSlidingWindow(nums []int, k int) []int {
	monoQueue := []int{} // store ind
	winSt := 0
	ans := []int{}
	for winEnd := 0; winEnd < len(nums); winEnd++ {
		for len(monoQueue) != 0 && nums[winEnd] > nums[monoQueue[len(monoQueue)-1]] {
			monoQueue = monoQueue[:len(monoQueue)-1]
		}
		monoQueue = append(monoQueue, winEnd)
		if winEnd < k-1 {
			continue
		}
		for monoQueue[0] < winSt {
			monoQueue = monoQueue[1:]
		}
		ans = append(ans, nums[monoQueue[0]])
		winSt++
	}
	return ans
}

// type MaxHeap [][2]int // <val, ind>

// func (h MaxHeap) Len() int           { return len(h) }
// func (h MaxHeap) Less(i, j int) bool { return h[i][0] > h[j][0] }
// func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
// func (h *MaxHeap) Push(x interface{}) {
// 	*h = append(*h, x.([2]int))
// }
// func (h *MaxHeap) Pop() interface{} {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[:n-1]
// 	return x
// }
// // priority queue method: time=O(nlogn),space=O(n)
// func maxSlidingWindow(nums []int, k int) []int {
// 	winSt := 0
// 	var mh MaxHeap
// 	heap.Init(&mh)
// 	ans := []int{}
// 	for winEnd := 0; winEnd < len(nums); winEnd++ {
// 		heap.Push(&mh, [2]int{nums[winEnd], winEnd})
// 		if winEnd < k-1 {
// 			continue
// 		}
// 		for mh[0][1] < winSt {
// 			heap.Pop(&mh)
// 		}
// 		ans = append(ans, mh[0][0])
// 		winSt++
// 	}
// 	return ans
// }
