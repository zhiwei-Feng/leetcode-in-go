package n0714

func maxProfit(prices []int, fee int) int {
	if len(prices) <= 0 {
		return 0
	}
	dp := make([][2]int, len(prices))
	dp[0][0] = -prices[0] - fee
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i]-fee) // 当天持有股票
		dp[i][1] = max(dp[i-1][0]+prices[i], dp[i-1][1])     // 当天不持有股票，即当天卖出)
	}
	return max(dp[len(prices)-1][0], dp[len(prices)-1][1])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
