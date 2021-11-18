package n0072

// dp[i][j] = minDistance(word1[:i], word2[:j])
// dp[i][j] = dp[i-1][j-1] if w1[i]==w2[j] else min(dp[i-1][j-1],dp[i][j-1],dp[i-1][j])+1
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	count := 0
	for i := 0; i < m; i++ {
		if word1[i] == word2[0] {
			count++
		}
		if count > 0 {
			dp[i][0] = i
		} else {
			dp[i][0] = i + 1
		}
	}
	count = 0
	for j := 0; j < n; j++ {
		if word1[0] == word2[j] {
			count++
		}
		if count > 0 {
			dp[0][j] = j
		} else {
			dp[0][j] = j + 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}

	return dp[m-1][n-1]
}

func min(x1, x2, x3 int) int {
	minI := x1
	if x2 < minI {
		minI = x2
	}
	if x3 < minI {
		minI = x3
	}
	return minI
}
