package n0005

import (
	"strings"
)

// 1. Manacher
func longestPalindrome(s string) string {
	// 插入#
	var build strings.Builder
	build.WriteString("#")
	for i := 0; i < len(s); i++ {
		build.WriteByte(s[i])
		build.WriteString("#")
	}
	s = build.String()

	// start
	dp_arm := make([]int, 0, len(s)) // 暂存前面位置的回文串臂长
	rightmost := -1                  // 当前最右位置
	j := -1                          // 最右位置对应的中心点位置
	start, end := 0, -1              // 最长回文串的开始和结束index
	for i := 0; i < len(s); i++ {
		cur_arm_len := 0 // 当前位置i对应的臂长
		if i <= rightmost {
			// 核心代码
			mirro := 2*j - i
			cur_arm_len = min(dp_arm[mirro], rightmost-i)
			cur_arm_len = expand_arm(s, i-cur_arm_len, i+cur_arm_len)
		} else {
			cur_arm_len = expand_arm(s, i, i)
		}

		dp_arm = append(dp_arm, cur_arm_len)

		// 保持rightmost和j的更新
		if cur_arm_len+i > rightmost {
			j = i
			rightmost = i + cur_arm_len
		}

		// 更新最长回文串
		if 2*cur_arm_len+1 > end-start+1 {
			start = i - cur_arm_len
			end = i + cur_arm_len
		}
	}

	// 处理#,并返回结果
	var ans strings.Builder
	for i := start; i <= end; i++ {
		if s[i] != '#' {
			ans.WriteByte(s[i])
		}
	}
	return ans.String()
}

func expand_arm(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return (right - left - 2) / 2
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// 2.center expand
func longestPalindrome_centerexpand(s string) string {
	n := len(s)
	start := 0
	end := 0
	for i := 0; i < n; i++ {
		len1 := expand(s, i, i)
		len2 := expand(s, i, i+1)
		len := max(len1, len2)
		if len > end-start+1 {
			start = i - (len-1)/2
			end = i + len/2
		}
	}

	return s[start : end+1]
}

func expand(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// 3. dp
func longestPalindrome_dp(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	ans := string(s[0])
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
		if i < len(s)-1 && s[i] == s[i+1] {
			dp[i][i+1] = true
			ans = s[i : i+2]
		}
	}

	for l := 3; l <= len(s); l++ {
		for i := 0; i <= len(s)-l; i++ {
			j := i + l - 1
			dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			if dp[i][j] && j-i+1 > len(ans) {
				ans = s[i : j+1]
			}
		}
	}

	return ans
}
