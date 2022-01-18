package offer0048

// dp做法
func lengthOfLongestSubstring(s string) int {
	recentIndex := make(map[byte]int)
	dp := make([]int, len(s))
	recentIndex[s[0]] = 0
	dp[0] = 1
	ans := 1
	for j := 1; j < len(s); j++ {
		i, ok := recentIndex[s[j]]
		if !ok {
			dp[j] = dp[j-1] + 1
			recentIndex[s[j]] = j
			if ans < dp[j] {
				ans = dp[j]
			}
			continue
		}
		// update index
		recentIndex[s[j]] = j

		if dp[j-1] < j-i {
			// 不在区间内
			dp[j] = dp[j-1] + 1
		} else {
			dp[j] = j - i
		}
		if ans < dp[j] {
			ans = dp[j]
		}
	}
	return ans
}

// 滑动窗口版本
func lengthOfLongestSubstring1(s string) int {
	winStart := 0
	winMap := make(map[byte]int)
	ans := 0
	for winEnd := 0; winEnd < len(s); winEnd++ {
		winMap[s[winEnd]]++

		for len(winMap) != winEnd-winStart+1 {
			winMap[s[winStart]]--
			if winMap[s[winStart]] == 0 {
				delete(winMap, s[winStart])
			}
			winStart++
		}

		if ans < len(winMap) {
			ans = len(winMap)
		}
	}
	return ans
}
