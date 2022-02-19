package n0739

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = n + 1
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; {
			if temperatures[j] > temperatures[i] {
				dp[i] = min(dp[i], j-i)
				break
			} else {
				j += dp[j]
			}
		}
	}
	for i := range dp {
		if dp[i] == n+1 {
			dp[i] = 0
		}
	}
	return dp
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
