package n1143

func minDistance(word1 string, word2 string) int {
	// first find the longest common subsequence
	m, n := len(word1), len(word2)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// init dp[0][j],dp[i][0]
	count := 0
	for i := 0; i < m; i++ {
		if word1[i] == word2[0] {
			count++
		}
		if count > 0 {
			dp[i][0] = 1
		}
	}
	count = 0
	for j := 0; j < n; j++ {
		if word1[0] == word2[j] {
			count++
		}
		if count > 0 {
			dp[0][j] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	// just reduced by length of word1,word2
	return m + n - 2*dp[m-1][n-1]
}

func max(x1, x2 int) int {
	if x1 > x2 {
		return x1
	}
	return x2
}
