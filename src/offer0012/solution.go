package offer0012

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	if m*n < len(word) {
		return false
	}

	var dfs func(i, j, count int) bool
	dfs = func(i, j, count int) bool {
		if count >= len(word) {
			return true
		}
		if i < 0 || j < 0 || i >= m || j >= n || board[i][j] != word[count] {
			return false
		}

		board[i][j] = '.'
		res := dfs(i-1, j, count+1) || dfs(i+1, j, count+1) || dfs(i, j-1, count+1) || dfs(i, j+1, count+1)
		board[i][j] = word[count]
		return res
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
