package n0072

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	// 将问题定义为
	// dp[i][j]表示word1前i+1个字符和word2前j+1个字符的编辑距离
	// 最终返回dp[m-1][n-1]
	// 其中转移过程，当word1[i-1]==word2[j-1]，直接由dp[i-1][j-1]进行转移
	// 不等于时，可以是由word1添加一个字符或者word2添加一个字符，或者两者都添加一个字符并执行替换操作，描述如下
	// dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1

	// 边界 dp[0][x], dp[x][0], 前者当word2[:x+1]中存在word1[0]时，只需要添加x个字符，后者同理
	sameExisted := false
	for i := 0; i < m; i++ {
		if word1[i] == word2[0] {
			sameExisted = true
		}
		if sameExisted {
			dp[i][0] = i
		} else {
			dp[i][0] = i + 1
		}
	}
	sameExisted = false
	for j := 0; j < n; j++ {
		if word1[0] == word2[j] {
			sameExisted = true
		}
		if sameExisted {
			dp[0][j] = j
		} else {
			dp[0][j] = j + 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}

	return dp[m-1][n-1]
}

func min(x1, x2, x3 int) int {
	minI := x1
	if x2 < minI {
		minI = x2
	}
	if x3 < minI {
		minI = x3
	}
	return minI
}
