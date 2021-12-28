package n0123

// buy1: 当前进行过一次买的最大值
// sell1: 当前进行过一次完整交易的最大值
// buy2: 当前进行一次交易且和第二次买的最大值
// sell2: 当前进行两次完整交易的最大值
func maxProfit(prices []int) int {
    buy1, sell1 := -prices[0], 0
    buy2, sell2 := -prices[0], 0
    for i := 1; i < len(prices); i++ {
        buy1 = max(buy1, -prices[i])
        sell1 = max(sell1, buy1+prices[i])
        buy2 = max(buy2, sell1-prices[i])
        sell2 = max(sell2, buy2+prices[i])
    }
    return sell2
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}