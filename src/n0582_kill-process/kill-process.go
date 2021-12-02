package n0582killprocess

func killProcess(pid []int, ppid []int, kill int) []int {
	parentToChild := make(map[int][]int)
	n := len(pid)
	for i := 0; i < n; i++ {
		id, p := pid[i], ppid[i]
		if p == 0 {
			continue
		}
		parentToChild[p] = append(parentToChild[p], id)
	}
	ans := []int{}
	queue := []int{}
	queue = append(queue, kill)
	for len(queue) > 0 {
		curKillProc := queue[0]
		queue = queue[1:]
		ans = append(ans, curKillProc)
		queue = append(queue, parentToChild[curKillProc]...)
	}
	return ans
}
