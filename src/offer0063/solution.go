package offer0063

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([]int, len(prices))
	dp[0] = 0
	hisMin := prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i] = max(dp[i-1], prices[i]-hisMin)
		hisMin = min(hisMin, prices[i])
	}
	return dp[len(prices)-1]
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
