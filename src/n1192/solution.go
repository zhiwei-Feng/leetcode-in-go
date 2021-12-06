package n1192

func criticalConnections(n int, connections [][]int) [][]int {
	ans := make([][]int, 0, len(connections))
	IDs := make([]int, n)
	edge := make([][]int, n)
	for i := 0; i < n; i++ {
		IDs[i] = -1
	}
	for _, conn := range connections {
		edge[conn[0]] = append(edge[conn[0]], conn[1])
		edge[conn[1]] = append(edge[conn[1]], conn[0])
	}

	var dfs func(int, int, int) int
	dfs = func(curNode int, from int, idShoudBe int) int {
		IDs[curNode] = idShoudBe
		for _, next := range edge[curNode] {
			if next == from {
				continue
			}
			if IDs[next] == -1 {
				// can visit
				IDs[curNode] = min(IDs[curNode], dfs(next, curNode, idShoudBe+1))
			} else {
				IDs[curNode] = min(IDs[next], IDs[curNode])
			}
		}
		if IDs[curNode] == idShoudBe && curNode != 0 {
			ans = append(ans, []int{from, curNode})
		}
		return IDs[curNode]
	}
	dfs(0, -1, 0)
	return ans
}

func min(x, y int) int {
	if x<y{
		return x
	}
	return y
}