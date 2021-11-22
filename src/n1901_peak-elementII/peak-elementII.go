package n1901

func findPeakGrid(mat [][]int) []int {
	m, _ := len(mat), len(mat[0])
	low, high := 0, m
	for low < high {
		mid := low + (high-low)/2
		var upMax int
		var downMax int
		var curMax int
		if mid == 0 {
			upMax = -1
		} else {
			_, upMax = maxOfRow(mat, mid-1)
		}
		if mid == m-1 {
			downMax = -1
		} else {
			_, downMax = maxOfRow(mat, mid+1)
		}
		canIdx, curMax := maxOfRow(mat, mid)
		if curMax >= upMax && curMax >= downMax {
			return []int{mid, canIdx}
		} else if upMax >= curMax && upMax >= downMax {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return []int{}
}

func maxOfRow(mat [][]int, row int) (int, int) {
	// return idx, value
	idx := 0
	for i := 1; i < len(mat[row]); i++ {
		if mat[row][i] > mat[row][idx] {
			idx = i
		}
	}

	return idx, mat[row][idx]
}
