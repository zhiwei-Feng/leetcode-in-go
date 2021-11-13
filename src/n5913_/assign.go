package n5913

import "sort"

func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(tasks)))
	sort.Sort(sort.Reverse(sort.IntSlice(workers)))
	m := len(tasks)
	n := len(workers)
	lo, hi := 0, min(m, n)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if success(tasks[m-mid:m], workers[:mid], pills, strength) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return lo
}

func success(tasks []int, workers []int, pills, strength int) bool {
	workerMap := make(map[int]int)
	for _, w := range workers {
		workerMap[w]++
	}
	for i := 0; i < len(tasks); i++ {
		if len(workerMap) == 0 {
			return false
		}

		// 抽取keys
		workerKeys := []int{}
		for k, _ := range workerMap {
			workerKeys = append(workerKeys, k)
		}
		sort.Ints(workerKeys)

		curWorker := workers[len(workers)-1]
		if curWorker >= tasks[i] {
			workerMap[curWorker]--
			if workerMap[curWorker] == 0 {
				delete(workerMap, curWorker)
			}
		} else {
			if pills == 0 {
				return false
			}
			ind := sort.Search(len(workerKeys), func(k int) bool {
				return workerKeys[k] >= tasks[i]-strength
			})
			if ind == len(workerKeys) {
				return false
			}
			pills--
			workerMap[workerKeys[ind]]--
			if workerMap[workerKeys[ind]] == 0 {
				delete(workerMap, workerKeys[ind])
			}
		}
	}
	return true
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
