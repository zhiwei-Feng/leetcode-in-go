package n0417

func pacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])
	directs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	ans := [][]int{}
	P := make([][]bool, m)
	A := make([][]bool, m)
	for i := 0; i < m; i++ {
		P[i] = make([]bool, n)
		A[i] = make([]bool, n)
	}

	bfs := func(visit [][]bool, isP bool) {
		queue := [][2]int{}
		if isP {
			for i := 0; i < m; i++ {
				queue = append(queue, [2]int{i, 0})
				visit[i][0] = true
			}
			for j := 1; j < n; j++ {
				queue = append(queue, [2]int{0, j})
				visit[0][j] = true
			}
		} else {
			for i := 0; i < m; i++ {
				queue = append(queue, [2]int{i, n - 1})
				visit[i][n - 1] = true
			}
			for j := 0; j < n-1; j++ {
				queue = append(queue, [2]int{m - 1, j})
				visit[m - 1][j] = true
			}
		}

		for len(queue) > 0 {
			levelNum := len(queue)
			for i := 0; i < levelNum; i++ {
				top := queue[0]
				queue = queue[1:]
				if P[top[0]][top[1]] && A[top[0]][top[1]] {
					ans = append(ans, []int{top[0], top[1]})
				}
				for _, d := range directs {
					newX, newY := top[0]+d[0], top[1]+d[1]
					if newX >= 0 && newX < m && newY >= 0 && newY < n {
						if !visit[newX][newY] && heights[newX][newY] >= heights[top[0]][top[1]] {
							queue = append(queue, [2]int{newX, newY})
							visit[newX][newY] = true
						}
					}
				}
			}
		}
	}

	bfs(P, true)
	bfs(A, false)
	return ans
}
