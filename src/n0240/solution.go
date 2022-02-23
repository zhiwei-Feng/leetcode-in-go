package n0240

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	sx, sy := 0, n-1
	for sx >= 0 && sy >= 0 && sx < m && sy < n {
		if target == matrix[sx][sy] {
			return true
		}
		if target > matrix[sx][sy] {
			sx++
		} else {
			sy--
		}
	}
	return false
}
