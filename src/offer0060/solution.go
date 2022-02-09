package offer0060

func dicesProbability(n int) []float64 {
	dp := make([]float64, 6)
	for i := range dp {
		dp[i] = 1.0 / 6
	}

	for i := 2; i < n; i++ {
		tmp := make([]float64, 5*i+1)
		for j := range dp {
			for k := 0; k < 6; k++ {
				tmp[j+k] += dp[j] / 6
			}
		}

		dp = tmp
	}
	return dp
}
