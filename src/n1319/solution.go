package n1319

// 并查集
func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}
	fa := make([]int, n)
	for i := 0; i < n; i++ {
		fa[i] = i
	}
	var find func(int) int
	var merge func(int, int)
	find = func(i int) int {
		if fa[i] != i {
			i = find(fa[i])
		}
		return i
	}
	count := n
	merge = func(x, y int) {
		x = find(x)
		y = find(y)
		if x == y {
			return
		}
		if x > y {
			x, y = y, x
		}
		fa[x] = y
		count--
	}

	for _, conn := range connections {
		merge(conn[0], conn[1])
	}

	return count - 1
}

func makeConnected_dfs(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}
	visit := make([]bool, n)
	graph := make([][]int, n)
	for _, conn := range connections {
		graph[conn[0]] = append(graph[conn[0]], conn[1])
		graph[conn[1]] = append(graph[conn[1]], conn[0])
	}

	var dfs func(int)
	dfs = func(cur int) {
		visit[cur] = true
		for _, node := range graph[cur] {
			if !visit[node] {
				dfs(node)
			}
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		if !visit[i] {
			dfs(i)
			ans++
		}
	}

	return ans - 1
}
