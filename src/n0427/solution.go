package n0427

type Node struct {
	Val, IsLeaf                                bool
	TopLeft, TopRight, BottomLeft, BottomRight *Node
}

func construct(grid [][]int) *Node {
	if len(grid) == 0 {
		return nil
	}
	allSame := true
	base := grid[0][0]
	n := len(grid)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == base {
				continue
			} else {
				allSame = false
				break
			}
		}
	}

	if allSame {
		return &Node{
			Val:    base == 1,
			IsLeaf: true,
		}
	} else {
		subRoot := &Node{Val: true, IsLeaf: false}
		subRoot.TopLeft = construct(subMatrix(grid, 0, 0, n/2, n/2))
		subRoot.TopRight = construct(subMatrix(grid, 0, n/2, n/2, n))
		subRoot.BottomLeft = construct(subMatrix(grid, n/2, 0, n, n/2))
		subRoot.BottomRight = construct(subMatrix(grid, n/2, n/2, n, n))
		return subRoot
	}
}

func subMatrix(grid [][]int, sx, sy, ex, ey int) [][]int {
	// sx <= x < ex
	res := make([][]int, ex-sx)
	for i := range res {
		res[i] = make([]int, ey-sy)
		copy(res[i], grid[sx+i][sy:ey])
	}
	return res
}
