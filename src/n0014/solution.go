package n0014

import "math"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	minLen := math.MaxInt64
	for _, v := range strs {
		if len(v) < minLen {
			minLen = len(v)
		}
	}
	l, r := 0, minLen
	for l < r {
		m := l + (r-l+1)/2
		if !match(strs, m) {
			r = m - 1
		} else {
			l = m
		}
	}
	return strs[0][:l]
}

func match(strs []string, m int) bool {
	// 匹配前m个字符是否一致
	for i := 0; i < m; i++ {
		base := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != base {
				return false
			}
		}
	}
	return true
}
