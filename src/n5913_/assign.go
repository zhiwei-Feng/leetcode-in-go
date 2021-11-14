package n5913

import "sort"

func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	sort.Ints(tasks)
	sort.Ints(workers)
	lo, hi := 0, min(len(tasks), len(workers))
	for lo < hi {
		// 这里不用lo+(hi-lo)/2是因为
		// 当出现仅有i,i+1的候选位置，且i位置是ok的时候，m会永远得到i，而不是进一步考察i+1
		m := lo + (hi-lo+1)/2
		task_ind := 0
		p := pills
		fail := false
		valid_tasks := []int{}
		for _, w := range workers[len(workers)-m:] {
			for task_ind < m && tasks[task_ind] <= w+strength {
				valid_tasks = append(valid_tasks, tasks[task_ind])
				task_ind++
			}
			if len(valid_tasks) == 0 {
				fail = true
				break
			}
			if valid_tasks[0] <= w {
				valid_tasks = valid_tasks[1:]
			} else {
				if p == 0 {
					fail = true
					break
				} else {
					p--
					valid_tasks = valid_tasks[:len(valid_tasks)-1]
				}
			}
		}
		if fail {
			hi = m - 1
		} else {
			lo = m
		}
	}
	return lo
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
