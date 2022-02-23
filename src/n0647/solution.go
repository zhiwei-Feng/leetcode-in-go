package n0647

import "strings"

func countSubstrings(s string) int {
	var build strings.Builder
	build.WriteString("$#")
	for i := range s {
		build.WriteByte(s[i])
		build.WriteByte('#')
	}
	build.WriteByte('!')
	t := build.String()
	n := len(t)
	f := make([]int, n)
	iMax, rMax, ans := 0, 0, 0
	for i := 1; i < n-1; i++ {
		if i <= rMax {
			f[i] = min(rMax-i+1, f[2*iMax-i])
		} else {
			f[i] = 1
		}
		// 中心扩展
		for t[i+f[i]] == t[i-f[i]] {
			f[i]++
		}
		if i+f[i]-1 > rMax {
			iMax = i
			rMax = i + f[i] - 1
		}
		ans += f[i] / 2
	}
	return ans
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func countSubstrings_extend(s string) int {
	if len(s) == 0 {
		return 0
	}
	ans := 0
	for i := 0; i < len(s); i++ {
		extend(s, i, i, &ans)
		if i < len(s)-1 {
			extend(s, i, i+1, &ans)
		}
	}
	return ans
}

func extend(s string, start, end int, ans *int) {
	for start >= 0 && end < len(s) && s[start] == s[end] {
		*ans++
		start--
		end++
	}
}

func countSubstrings_(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	dp := make([][]bool, n) // i..j 是否为回文串
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	ans := n
	for i := 0; i < n; i++ {
		dp[i][i] = true
		if i > 0 && s[i-1] == s[i] {
			dp[i-1][i] = true
			ans++
		}
	}
	for k := 3; k <= n; k++ {
		for i := 0; i <= n-k; i++ {
			dp[i][i+k-1] = dp[i+1][i+k-2] && s[i] == s[i+k-1]
			if dp[i][i+k-1] {
				ans++
			}
		}
	}
	return ans
}
