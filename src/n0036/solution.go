package n0036

func isValidSudoku(board [][]byte) bool {
	rows, columns := [9][9]int{}, [9][9]int{}
	subboxs := [3][3][9]int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			number := board[i][j] - '0' - 1
			rows[i][number]++
			columns[j][number]++
			subboxs[i/3][j/3][number]++
			if rows[i][number] > 1 || columns[j][number] > 1 || subboxs[i/3][j/3][number] > 1 {
				return false
			}
		}
	}
	return true
}
