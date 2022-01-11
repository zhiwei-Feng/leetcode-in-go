package n1036

var directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

type Pos struct {
	X, Y int
}

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	n := len(blocked)
	if n < 2 {
		return true
	}

	const (
		block        = -1
		valid        = 0
		found        = 1
		boundary int = 1e6
	)
	// blockCircleLimit := n * (n - 1) / 2
	blockedMap := make(map[Pos]bool)
	for _, v := range blocked {
		blockedMap[Pos{v[0], v[1]}] = true
	}

	check := func(start, finish []int) int {
		sx, sy := start[0], start[1]
		fx, fy := finish[0], finish[1]
		breakdown := n * (n - 1) / 2

		q := []Pos{{sx, sy}}
		vis := map[Pos]bool{Pos{sx, sy}: true}
		for len(q) != 0 && breakdown > 0 {
			cur := q[0]
			q = q[1:]
			for _, d := range directions {
				newx, newy := cur.X+d[0], cur.Y+d[1]
				nPos := Pos{newx, newy}
				if newx >= 0 && newx < boundary && newy >= 0 && newy < boundary && !blockedMap[nPos] && !vis[nPos] {
					if newx == fx && newy == fy {
						return found
					}
					breakdown--
					vis[nPos] = true
					q = append(q, nPos)
				}
			}
		}

		if breakdown > 0 {
			return block
		}
		return valid
	}

	res := check(source, target)
	if res == found {
		return true
	} else if res == block {
		return false
	} else {
		rev := check(target, source)
		if rev == block {
			return false
		}
		return true
	}
}
