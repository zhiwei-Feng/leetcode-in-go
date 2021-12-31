package n1136

func minimumSemesters(n int, relations [][]int) int {
	graph := make([][]int, n+1)
	inDegree := make([]int, n+1)
	for _, conn := range relations {
		head, tail := conn[0], conn[1]
		graph[head] = append(graph[head], tail)
		inDegree[tail]++
	}

	queue := []int{}
	for i := 1; i <= n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	depth := 0
	for len(queue) != 0 {
		depth++
		level := len(queue)
		for i := 0; i < level; i++ {
			top := queue[0]
			queue = queue[1:]
			for _, v := range graph[top] {
				inDegree[v]--
				if inDegree[v] == 0 {
					queue = append(queue, v)
				}
			}
		}
	}

	allDone := true
	for i := 1; i <= n; i++ {
		if inDegree[i] != 0 {
			allDone = false
			break
		}
	}
	if allDone {
		return depth
	} else {
		return -1
	}
}
