package n0322

import (
	"math"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = math.MaxInt16 // init dp
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			// dp[i] = min(dp[i-coin[0]], dp[i-coin[1]]...) + 1
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount]==math.MaxInt16 {
		return -1
	}
	return dp[amount]
}

func min(x1, x2 int) int {
	if x1 < x2 {
		return x1
	}
	return x2
}

func coinChange_memory_search(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}
	count := make([]int, amount+1)
	return dp(coins, amount, count)
}

func dp(coins []int, remain int, count []int) int {
	if remain < 0 {
		return -1
	}
	if remain == 0 {
		return 0
	}
	if count[remain] != 0 {
		return count[remain]
	}

	tmpMin := math.MaxInt16
	for _, coin := range coins {
		res := dp(coins, remain-coin, count)
		if res >= 0 && res < tmpMin {
			tmpMin = res + 1
		}
	}
	if tmpMin != math.MaxInt16 {
		count[remain] = tmpMin
	} else {
		count[remain] = -1
	}

	return count[remain]
}
