package n0347

import "math/rand"

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	elems := []Element{}
	for k, v := range m {
		elems = append(elems, Element{Val: k, Freq: v})
	}
	i, j := 0, len(elems)-1
	for i < j {
		pivotIdx := partition(elems, i, j)
		if pivotIdx == k-1 {
			i = pivotIdx
			break
		} else if pivotIdx < k-1 {
			i = pivotIdx + 1
		} else {
			j = pivotIdx - 1
		}
	}
	ans := []int{}
	for k := 0; k <= i; k++ {
		ans = append(ans, elems[k].Val)
	}
	return ans
}

func partition(elems []Element, start, end int) int {
	randIdx := rand.Intn(end-start+1) + start
	elems[start], elems[randIdx] = elems[randIdx], elems[start]
	pivot := elems[start].Freq
	leftIdx := start + 1
	for i := leftIdx; i <= end; i++ {
		if elems[i].Freq > pivot {
			elems[i], elems[leftIdx] = elems[leftIdx], elems[i]
			leftIdx++
		}
	}
	elems[start], elems[leftIdx-1] = elems[leftIdx-1], elems[start]
	return leftIdx - 1
}

type Element struct {
	Val, Freq int
}
