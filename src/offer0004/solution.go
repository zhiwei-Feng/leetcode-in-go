package offer0004

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	sx, sy := len(matrix)-1, 0
	for sx >= 0 && sy < len(matrix[0]) {
		if matrix[sx][sy] == target {
			return true
		}

		if target > matrix[sx][sy] {
			sy++
		} else {
			sx--
		}
	}

	return false
}
