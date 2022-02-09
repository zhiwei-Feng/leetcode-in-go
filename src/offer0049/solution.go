package offer0049

func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		num2, num3, num5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(min(num2, num3), num5)
		if dp[i] == num2 {
			p2++
		}
		if dp[i] == num3 {
			p3++
		}
		if dp[i] == num5 {
			p5++
		}
	}
	return dp[n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
