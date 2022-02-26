package n0438

func findAnagrams(s string, p string) []int {
	k := len(p)
	winMap := make(map[byte]int)
	for i := range p {
		winMap[p[i]]++
	}
	winStart := 0
	indexList := []int{}
	winMatch := 0
	for winEnd := 0; winEnd < len(s); winEnd++ {
		newByte := s[winEnd]
		if _, ok := winMap[newByte]; ok {
			winMap[newByte]--
			if winMap[newByte] == 0 {
				winMatch++
			}
		}
		if winEnd < k-1 {
			continue
		}
		if winMatch == len(winMap) {
			indexList = append(indexList, winStart)
		}
		delByte := s[winStart]
		if _, ok := winMap[delByte]; ok {
			if winMap[delByte] == 0 {
				winMatch--
			}
			winMap[delByte]++
		}
		winStart++
	}
	return indexList
}
