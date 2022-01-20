package offer0058_1

import "strings"

func reverseWords(s string) string {
	var builder strings.Builder
	winEnd := len(s)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if winEnd-i > 1 {
				builder.WriteString(s[i+1 : winEnd])
				builder.WriteString(" ")
			}
			winEnd = i
		}
	}
	// 处理最后一个
	if winEnd > 0 {
		builder.WriteString(s[:winEnd])
		builder.WriteString(" ")
	}

	ans := builder.String()
	if len(ans) == 0 {
		return ""
	}
	return ans[:len(ans)-1]
}
