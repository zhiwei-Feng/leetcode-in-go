package n0279

import "math"

func numSquares(n int) int {
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minn := math.MaxInt32
		for j:=1;j*j<=i;j++ {
			minn = min(minn, f[i-j*j]+1)
		}
		f[i]=minn
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
