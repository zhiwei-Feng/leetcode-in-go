package n1765

func highestPeak(isWater [][]int) [][]int {
	next := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(isWater), len(isWater[0])
	height := make([][]int, m)
	for i, _ := range height {
		height[i] = make([]int, n)
		for j, _ := range height[i] {
			height[i][j] = -1
		}
	}
	type pair struct {
		x, y int
	}

	q := []pair{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				q = append(q, pair{i, j})
				height[i][j] = 0
			}
		}
	}

	for len(q) != 0 {
		curx, cury := q[0].x, q[0].y
		q = q[1:]
		curh := height[curx][cury]
		for _, d := range next {
			newx, newy := d[0]+curx, d[1]+cury
			if newx < 0 || newx >= m || newy < 0 || newy >= n || height[newx][newy] != -1 {
				continue
			}
			height[newx][newy] = curh + 1
			q = append(q, pair{newx, newy})
		}
	}

	return height
}
