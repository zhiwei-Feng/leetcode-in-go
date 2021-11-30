package n0340

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	winSt := 0
	winMap := make(map[byte]int)
	ans := 0
	for winEn := 0; winEn < len(s); winEn++ {
		winMap[s[winEn]]++
		for len(winMap) > k {
			delByte := s[winSt]
			winMap[delByte]--
			if winMap[delByte] == 0 {
				delete(winMap, delByte)
			}
			winSt++
		}
		if winEn-winSt+1 > ans {
			ans = winEn - winSt + 1
		}
	}
	return ans
}
