package n1368

func minCost(grid [][]int) int {
	direction := [][2]int{{0, 0}, {0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	if len(grid) == 1 && len(grid[0]) == 1 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	dist := make([][]int, m)
	for i := 0; i < m; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = 2<<16 - 1
		}
	}

	// begin bfs
	queue := []Pair{}
	queue = append(queue, Pair{0, 0})
	dist[0][0] = 0
	for len(queue) > 0 {
		x, y := queue[0].X, queue[0].Y
		queue = queue[1:]
		for i := 1; i <=4; i++ {
			newX, newY := x+direction[i][0], y+direction[i][1]
			if newX < 0 || newX >= m || newY < 0 || newY >= n {
				continue
			}
			tmpDist := dist[x][y]
			if i != grid[x][y] {
				tmpDist++
			}
			if tmpDist < dist[newX][newY] {
				dist[newX][newY] = tmpDist
				queue = append(queue, Pair{newX, newY})
			}
		}
	}
	return dist[m-1][n-1]
}

type Pair struct {
	X int
	Y int
}
