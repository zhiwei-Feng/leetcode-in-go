package bytedance0004

// 输出有向无环图的所有拓扑序列
func solve(n int, edges [][]int) [][]int {
	ans := [][]int{}
	inDegree := make([]int, n)
	visit := make([]bool, n)
	graph := make([][]bool, n)
	for i := range graph {
		graph[i] = make([]bool, n)
	}
	for _, e := range edges {
		v, w := e[0], e[1]
		graph[v][w] = true
		inDegree[w]++
	}

	cur := []int{} // 存储当前拓扑排序
	var dfs func()
	dfs = func() {
		if len(cur) == n {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			ans = append(ans, tmp)
			return
		}

		for i := 0; i < n; i++ {
			if !visit[i] && inDegree[i] == 0 {
				cur = append(cur, i)
				visit[i] = true
				for j := 0; j < n; j++ {
					if graph[i][j] {
						inDegree[j]--
					}
				}
				dfs()
				visit[i] = false
				cur = cur[:len(cur)-1]
				for j := 0; j < n; j++ {
					if graph[i][j] {
						inDegree[j]++
					}
				}
			}
		}
	}
	dfs()
	return ans
}
