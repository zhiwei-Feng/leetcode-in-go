package n0131

func partition(s string) [][]string {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
		if i+1 < n && s[i] == s[i+1] {
			dp[i][i+1] = true
		}
	}

	for k := 3; k <= n; k++ {
		for i := 0; i+k-1 < n; i++ {
			j := i + k - 1
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
			}
		}
	}

	ans := [][]string{}
	tmp := []string{}
	var backtrace func(pos int)
	backtrace = func(pos int) {
		if pos >= n {
			copyTmp := make([]string, len(tmp))
			copy(copyTmp, tmp)
			ans = append(ans, copyTmp)
			return
		}

		for k := 1; k <= n; k++ {
			i := pos
			if j := i + k - 1; j < n {
				if dp[i][j] {
					tmp = append(tmp, s[i:j+1])
					backtrace(j + 1)
					tmp = tmp[:len(tmp)-1]
				}
			}
		}

	}

	backtrace(0)
	return ans
}
