package n0486

func PredictTheWinner(nums []int) bool {
	n := len(nums)
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
		f[i][i] = nums[i]
		if i+1 < n {
			f[i][i+1] = max(nums[i], nums[i+1])
		}
	}

	//total
	total := 0
	for _, v := range nums {
		total += v
	}

	for k := 3; k <= n; k++ {
		for i := 0; i+k-1 < n; i++ {
			j := i + k - 1
			// f[i][j]
			var t1, t2 int
			// choose i
			t1 = nums[i] + min(f[i+1][j-1], f[i+2][j])
			// choose j
			t2 = nums[j] + min(f[i][j-2], f[i+1][j-1])

			f[i][j] = max(t1, t2)
		}
	}

	return f[0][n-1]*2 >= total
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
