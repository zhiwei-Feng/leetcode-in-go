package n0329

var directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	outDegree := make([][]int, m)
	for i := 0; i < m; i++ {
		outDegree[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for _, d := range directions {
				newX, newY := i+d[0], j+d[1]
				if newX >= 0 && newX < m && newY >= 0 && newY < n && matrix[i][j] < matrix[newX][newY] {
					outDegree[i][j]++
				}
			}
		}
	}

	queue := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if outDegree[i][j] == 0 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	ans := 0
	for len(queue) != 0 {
		ans++
		level := len(queue)
		for i := 0; i < level; i++ {
			curX, curY := queue[0][0], queue[0][1]
			queue = queue[1:]
			for _, d := range directions {
				newX, newY := curX+d[0], curY+d[1]
				if newX >= 0 && newY >= 0 && newX < m && newY < n && matrix[curX][curY] > matrix[newX][newY] {
					outDegree[newX][newY]--
					if outDegree[newX][newY] == 0 {
						queue = append(queue, []int{newX, newY})
					}
				}
			}
		}
	}

	return ans
}
