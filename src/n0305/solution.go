package n0305

func numIslands2(m int, n int, positions [][]int) []int {
	finder := Constructor(m * n)
	vis := make([]bool, m*n)
	ans := []int{}
	directs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, pos := range positions {
		currx, curry := pos[0], pos[1]
		index := currx*n + curry
		if !vis[index] {
			finder.addCount()
			vis[index] = true
			for _, d := range directs {
				newx, newy := currx+d[0], curry+d[1]
				newIndex := newx*n + newy
				if newx >= 0 && newy >= 0 && newx < m && newy < n && vis[newIndex] && !finder.isConnected(index, newIndex) {
					finder.union(index, newIndex)
				}
			}
		}
		ans = append(ans, finder.Count)
	}
	return ans
}

type UnionFind struct {
	Parents []int
	Count   int
}

func Constructor(N int) UnionFind {
	finder := UnionFind{
		Parents: make([]int, N),
		Count:   0,
	}
	for i := range finder.Parents {
		finder.Parents[i] = i
	}
	return finder
}

func (f *UnionFind) find(x int) int {
	for f.Parents[x] != x {
		x = f.find(f.Parents[x])
	}
	return f.Parents[x]
}

func (f *UnionFind) union(x, y int) {
	rootx := f.find(x)
	rooty := f.find(y)
	if rootx == rooty {
		return
	}
	f.Count--
	f.Parents[rootx] = rooty
}

func (f *UnionFind) isConnected(x, y int) bool {
	return f.find(x) == f.find(y)
}

func (f *UnionFind) addCount() {
	f.Count++
}
