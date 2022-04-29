package offerii0013

type NumMatrix struct {
	Sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 {
		return NumMatrix{}
	}
	m, n := len(matrix), len(matrix[0])
	sums := make([][]int, m+1)
	for i := range sums {
		sums[i] = make([]int, n+1)
		if i == 0 {
			continue
		}
		for j := 1; j <= n; j++ {
			sums[i][j] = sums[i-1][j] + sums[i][j-1] - sums[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return NumMatrix{Sums: sums}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.Sums[row2+1][col2+1] - this.Sums[row2+1][col1] - this.Sums[row1][col2+1] + this.Sums[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
