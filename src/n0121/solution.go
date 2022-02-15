package n0121

func maxProfit(prices []int) int {
	buyMin := prices[0]
	ans := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-buyMin > ans {
			ans = prices[i] - buyMin
		}
		if prices[i] < buyMin {
			buyMin = prices[i]
		}
	}
	return ans
}
