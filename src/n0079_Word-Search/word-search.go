package n0079

func exist(board [][]byte, word string) bool {
	var (
		m   = len(board)
		n   = len(board[0])
		k   = len(word)
		dfs func(int, int, int) bool
		dxy = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		vis = make([][]bool, m) // mark visisted status
	)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}

	dfs = func(x, y int, pos int) bool {
		if board[x][y]!=word[pos]{
			return false
		}
		if pos==k-1{
			return true
		}
		vis[x][y]=true
		defer func(){vis[x][y]=false}()
		for _,d:=range dxy{
			newX,newY:=x+d[0],y+d[1]
			if newX>=0&&newY>=0&&newX<m&&newY<n&&!vis[newX][newY]&&dfs(newX,newY,pos+1){
				return true
			}
		}
		return false
	}

	// entrance
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}

func exist_v1(board [][]byte, word string) bool {
	var (
		m   = len(board)
		n   = len(board[0])
		k   = len(word)
		dfs func(int, int, int)
		ans = false
		dxy = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		vis = make([][]bool, m) // mark visisted status
	)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}

	dfs = func(x, y int, pos int) {
		if pos == k {
			ans = true
			return
		}

		vis[x][y] = true
		defer func() { vis[x][y] = false }()
		for _, d := range dxy {
			newX, newY := x+d[0], y+d[1]
			if newX >= 0 && newX < m && newY >= 0 && newY < n && !vis[newX][newY] && board[newX][newY] == word[pos] {
				dfs(newX, newY, pos+1)
			}
		}

	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				dfs(i, j, 1)
				if ans {
					return ans
				}
			}
		}
	}

	return ans
}
