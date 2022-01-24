package offer0013

var dirs = [][2]int{{0, 1}, {1, 0}}

func movingCount(m int, n int, k int) int {
	type pos struct{ x, y int }
	q := []pos{{0, 0}}
	ans := 1
	visit := map[pos]bool{pos{0, 0}: true}
	for len(q) != 0 {
		x, y := q[0].x, q[0].y
		q = q[1:]
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			npos := pos{nx, ny}
			if nx < 0 || ny < 0 || nx >= m || ny >= n || !check(nx, ny, k) {
				continue
			}
			if visit[npos] {
				continue
			}
			visit[npos] = true
			ans++
			q = append(q, npos)
		}
	}
	return ans
}

func check(i, j, k int) bool {
	count := 0
	for i > 0 {
		count += i % 10
		i /= 10
	}
	for j > 0 {
		count += j % 10
		j /= 10
	}
	return count <= k
}
