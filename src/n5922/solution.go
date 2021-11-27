package n5922

func countWords(words1 []string, words2 []string) int {
	mapCount := make(map[string]int)
	ans := 0
	for _, word := range words1 {
		mapCount[word]++
	}

	for _, word := range words2 {
		if v, ok := mapCount[word]; ok {
			if v>1 {
				delete(mapCount, word)
				continue
			}
			mapCount[word]--
		}
	}

	for _, v := range mapCount {
		if v == 0 {
			ans++
		}
	}

	return ans
}
