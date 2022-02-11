package n0003

func lengthOfLongestSubstring(s string) (ans int) {
	winStart := 0
	winMap := make(map[byte]int)
	for winEnd := 0; winEnd < len(s); winEnd++ {
		winMap[s[winEnd]]++
		for len(winMap) != winEnd-winStart+1 {
			delChar := s[winStart]
			winMap[delChar]--
			if winMap[delChar] == 0 {
				delete(winMap, delChar)
			}
			winStart++
		}
		if len(winMap) > ans {
			ans = len(winMap)
		}
	}
	return ans
}
