package n0048

func rotate(matrix [][]int) {
	// 上下翻转
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-i-1] = matrix[n-i-1], matrix[i]
	}

	// 对角线翻转
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

// func rotate(matrix [][]int) {
// 	n := len(matrix)
// 	rotNum := n / 2
// 	for i := 0; i < rotNum; i++ {
// 		sx, ex, sy, ey := i, n-1-i, i, n-1-i
// 		for k := 0; k < ex-sx; k++ {
// 			rotateLevel(sx, ex, sy, ey, matrix)
// 		}
// 	}
// }

// func rotateLevel(sx, ex, sy, ey int, matrix [][]int) {
// 	prev := matrix[sx][sy]

// 	for i := sy + 1; i <= ey; i++ {
// 		prev, matrix[sx][i] = matrix[sx][i], prev
// 	}
// 	for i := sx + 1; i <= ex; i++ {
// 		prev, matrix[i][ey] = matrix[i][ey], prev
// 	}
// 	for i := ey - 1; i >= sy; i-- {
// 		prev, matrix[ex][i] = matrix[ex][i], prev
// 	}
// 	for i := ex - 1; i >= sx; i-- {
// 		prev, matrix[i][sy] = matrix[i][sy], prev
// 	}
// }
