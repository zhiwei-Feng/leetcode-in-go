package n0279

import "math"

func numSquares(n int) int {
	coins := make([]int, 0)
	for j := 1; j*j <= n; j++ {
		coins = append(coins, j*j)
	}
	f := make([][]int, len(coins))
	for i := 0; i < len(coins); i++ {
		f[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			f[i][j] = math.MaxInt32
		}
	}
	for i := 0; i < len(coins); i++ {
		f[i][0] = 0
	}
	for j := 1; j <= n; j++ {
		if j == coins[0] {
			f[0][j] = 1
		}
	}
	for i := 1; i < len(coins); i++ {
		for j := 0; j <= n; j++ {
			f[i][j] = f[i-1][j]
			if j >= coins[i] {
				f[i][j] = min(f[i][j], f[i][j-coins[i]]+1)
			}
		}
	}
	return f[len(coins)-1][n]
}

func numSquares_1d(n int) int {
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minn := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			minn = min(minn, f[i-j*j]+1)
		}
		f[i] = minn
	}
	if f[n] == math.MaxInt16 {
		return -1
	}
	return f[n]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
