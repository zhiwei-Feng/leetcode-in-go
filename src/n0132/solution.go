package n0132

import "math"

func minCut(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
		if i+1 < n && s[i] == s[i+1] {
			dp[i][i+1] = true
		}
	}

	for k := 3; k <= n; k++ {
		for i := 0; i+k-1 < n; i++ {
			j := i + k - 1
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
			}
		}
	}

	if dp[0][n-1] {
		return 0
	}

	f := make([]int, n)
	for i, _ := range f {
		if dp[0][i] {
			continue
		}

		f[i] = math.MaxInt32
		for j := 0; j < i; j++ {
			if dp[j+1][i] && f[j]+1 < f[i] {
				f[i] = f[j] + 1
			}
		}
	}
	return f[n-1]
}
