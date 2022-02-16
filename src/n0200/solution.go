package n0200

type pos struct{ x, y int }

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	ans := 0
	bfs := func(stx, sty int) {
		q := []pos{{stx, sty}}
		grid[stx][sty] = '0'
		for len(q) != 0 {
			curx, cury := q[0].x, q[0].y
			q = q[1:]
			for _, d := range dirs {
				newx, newy := curx+d[0], cury+d[1]
				if newx < 0 || newy < 0 || newx >= m || newy >= n || grid[newx][newy] == '0' {
					continue
				}
				grid[newx][newy] = '0'
				q = append(q, pos{newx, newy})
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				ans++
				bfs(i, j)
			}
		}
	}

	return ans
}
