package n0049

import "sort"

func groupAnagrams(strs []string) [][]string {
	group := make(map[string][]string)
	for _, v := range strs {
		byteArr := []byte(v)
		sort.SliceStable(byteArr, func(i, j int) bool { return byteArr[i] < byteArr[j] })
		sortedStr := string(byteArr)
		group[sortedStr] = append(group[sortedStr], v)
	}
	ans := [][]string{}
	for _, v := range group {
		ans = append(ans, v)
	}
	return ans
}
