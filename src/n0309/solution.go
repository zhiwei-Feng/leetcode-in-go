package n0309

// dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
// dp[i][1] = dp[i-1][0]+prices[i]
// dp[i][2] = max(dp[i-1][1], dp[i-1][2])
func maxProfit(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}
	dp := make([][3]int, len(prices))
	dp[0][0] = -prices[0]
	dp[0][1] = -1
	dp[0][2] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i]) // 当天持有股票
		dp[i][1] = dp[i-1][0] + prices[i]                // 当天不持有股票，且处于冷冻期(即当天卖出)
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])           // 当天不持有，不处于冷冻期(之前卖出)
	}
	return max(dp[len(prices)-1][1], dp[len(prices)-1][2])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
