package n0902

import (
	"strconv"
)

func atMostNGivenDigitSet(digits []string, n int) int {
	s := strconv.Itoa(n)
	K := len(s)
	// for example, n := 1234
	// dp[0] -> count of num <= n(1234) when length of num equal to 4
	// dp[1] -> count of num <= n[1:](234) when length of num equal to 3
	// ..
	dp := make([]int, K+1)
	dp[K] = 1

	for i := K - 1; i >= 0; i-- {
		Si := int(s[i] - '0')
		for _, d := range digits {
			val := int(d[0] - '0')
			if val < Si {
				dp[i] += pow(len(digits), K-i-1)
			} else if val == Si {
				dp[i] += dp[i+1]
			}
		}
	}

	for i := 1; i < K; i++ {
		dp[0] += pow(len(digits), i)
	}
	return dp[0]
}

func pow(base, n int) int {
	ans := 1
	add := base
	for n > 0 {
		if n&1 == 1 {
			ans = ans * add
		}
		add *= add
		n >>= 1
	}
	return ans
}
