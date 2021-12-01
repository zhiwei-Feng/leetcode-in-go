package n0076

func minWindow(s string, t string) string {
	winSt := 0
	wordMap := make(map[byte]int)
	wordCount := 0
	for i, _ := range t {
		wordMap[t[i]]++
	}
	ans := ""
	for winEnd := 0; winEnd < len(s); winEnd++ {
		if _, ok := wordMap[s[winEnd]]; ok {
			wordMap[s[winEnd]]--
			if wordMap[s[winEnd]] == 0 {
				wordCount++
			}
		}

		for wordCount == len(wordMap) {
			if ans == "" || len(ans) > winEnd-winSt+1 {
				ans = s[winSt : winEnd+1]
			}
			if _, ok := wordMap[s[winSt]]; ok {
				if wordMap[s[winSt]] == 0 {
					wordCount--
				}
				wordMap[s[winSt]]++
			}
			winSt++
		}
	}
	return ans
}
