package n0139

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)) // dp[i]表示[0..<=i]能否拆分
	wordCount := make(map[string]bool)
	for _, v:= range wordDict {
		wordCount[v]=true
	}
	if _, ok := wordCount[string(s[0])];ok{
		dp[0]=true
	}
	for i:=1;i<len(s);i++{
		if wordCount[s[:i+1]]{
			dp[i]=true
			continue
		}
		for j:=0;j<i;j++{
			// 切分，以[0..=j]为前一部分，[j+1..=i]为后一份
			last := s[j+1:i+1]
			if dp[j] {
				if _, ok:=wordCount[last];ok{
					dp[i]=true
					break
				}
			}
		}
	}
	return dp[len(s)-1]
}

func wordBreak_(s string, wordDict []string) bool {
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
