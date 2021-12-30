package n0310

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	graph := make(map[int][]int)
	inCount := make([]int, n)
	for _, conn := range edges {
		inCount[conn[0]]++
		inCount[conn[1]]++
		graph[conn[0]] = append(graph[conn[0]], conn[1])
		graph[conn[1]] = append(graph[conn[1]], conn[0])
	}

	var res []int
	q := []int{}
	for i, v := range inCount {
		if v == 1 {
			q = append(q, i)
		}
	}

	for len(q) != 0 {
		res = make([]int, 0)
		level := len(q)
		for k := 0; k < level; k++ {
			qt := q[0]
			q = q[1:]
			res = append(res, qt)
			for _, v := range graph[qt] {
				inCount[v]--
				if inCount[v] == 1 {
					q = append(q, v)
				}
			}
		}
	}

	return res
}
