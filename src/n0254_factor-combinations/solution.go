package n0254


func getFactors(n int) [][]int {
	return dfs(n, 2)
}

func dfs(n int, st int) [][]int {
	res := [][]int{}
	for i := st; i*i <= n; i++ {
		if n%i == 0 {
			res = append(res, []int{n / i, i})
			for _, v := range dfs(n/i, i) {
				v = append(v, i)
				res = append(res, v)
			}
		}
	}
	return res
}
