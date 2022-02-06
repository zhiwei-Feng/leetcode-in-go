package offer0029

import "math"

type pos struct{ x, y int }

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	directs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	curDir := 0
	start := pos{0, 0}
	ans := []int{}
	q := []pos{start}
	for len(q) != 0 {
		curx, cury := q[0].x, q[0].y
		q = q[1:]
		ans = append(ans, matrix[curx][cury])
		matrix[curx][cury] = math.MinInt32
		for i := 0; i < 4; i++ {
			newx, newy := curx+directs[curDir][0], cury+directs[curDir][1]
			if newx < 0 || newy < 0 || newx >= len(matrix) || newy >= len(matrix[0]) || matrix[newx][newy] == math.MinInt32 {
				curDir = (curDir + 1) % 4
				continue
			} else {
				q = append(q, pos{newx, newy})
				break
			}
		}
	}
	return ans
}
