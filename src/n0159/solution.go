package n0159

func lengthOfLongestSubstringTwoDistinct(s string) int {
	winSt := 0
	winMap := make(map[byte]int)
	ans := 0
	for winEn := 0; winEn < len(s); winEn++ {
		winMap[s[winEn]]++
		for len(winMap) > 2 {
			delByte := s[winSt]
			winMap[delByte]--
			if winMap[delByte] == 0 {
				delete(winMap, delByte)
			}
			winSt++
		}
		ans = max(ans, winEn-winSt+1)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
