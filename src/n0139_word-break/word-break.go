package n0139

func wordBreak(s string, wordDict []string) bool {
	count := make(map[string]int)
	minLen := 1<<7-1
	maxLen := 0
	for _, word := range wordDict {
		count[word]++
		if len(word) > maxLen {
			maxLen = len(word)
		}
		if len(word) < minLen {
			minLen = len(word)
		}
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := minLen; i <= len(s); i++ {
		for j := i - minLen; j >= 0 && j >= i-maxLen; j-- {
			_, ok := count[s[j:i]]
			dp[i] = dp[i] || (dp[j] && ok)
		}
	}

	return dp[len(s)]
}
