package n2022

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}
	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
		start := i * n
		end := (i + 1) * n
		copy(ans[i], original[start:end])
	}

	return ans
}
