package n0884

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	words1 := strings.Split(s1, " ")
	words2 := strings.Split(s2, " ")
	countmap := make(map[string]int)
	for _, w := range words1 {
		countmap[w]++
	}
	for _, w := range words2 {
		countmap[w]++
	}

	ans := []string{}
	for k, v := range countmap {
		if v==1 {
			ans = append(ans, k)
		}
	}
	return ans
}
