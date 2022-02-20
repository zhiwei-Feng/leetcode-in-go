package n0076

func minWindow(s string, t string) string {
	ans := ""
	winSt := 0
	winMap := make(map[byte]int)
	winMatch := 0
	for i := range t {
		winMap[t[i]]++
	}

	for winEnd := 0; winEnd < len(s); winEnd++ {
		curByte := s[winEnd]
		if _, ok := winMap[curByte]; ok {
			winMap[curByte]--
			if winMap[curByte] == 0 {
				winMatch++
			}
		}
		if winMatch < len(winMap) {
			continue
		}
		for winMatch == len(winMap) {
			// update ans
			if ans == "" || (winEnd-winSt+1) < len(ans) {
				ans = s[winSt : winEnd+1]
			}
			delByte := s[winSt]
			if _, ok := winMap[delByte]; ok {
				if winMap[delByte] == 0 {
					winMatch--
				}
				winMap[delByte]++
			}
			winSt++
		}
	}
	return ans
}

func minWindow_(s string, t string) string {
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
