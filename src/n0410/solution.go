package n0410

import "math"

// 1. dp[i][j]=MAXVAL if i<j
// 2. dp[0][0] = 0
// 3. dp[i][j] = min{max(dp[k][j-1], sub(k+1, i))} for k:=0;k<i
func splitArray(nums []int, m int) int {
	n := len(nums)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		for p := 0; p < len(dp[i]); p++ {
			dp[i][p] = math.MaxInt32
		}
	}
	sub := make([]int, n+1)
	for i := 0; i < n; i++ {
		sub[i+1] = sub[i] + nums[i]
	}
	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		for j := 1; j <= min(i, m); j++ {
			for k := 0; k < i; k++ {
				dp[i][j] = min(max(dp[k][j-1], sub[i]-sub[k]), dp[i][j])
			}
		}
	}

	return dp[n][m]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
