package n0394

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	var deBuilder strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			deBuilder.WriteByte(s[i])
		} else if s[i] >= '0' && s[i] <= '9' {
			// 0 is invalid
			// find entire integer
			number, stringStart := findNumber(i, s)
			end, res := decode(stringStart, s, number)
			deBuilder.WriteString(res)
			i = end
		}
	}
	return deBuilder.String()
}

func findNumber(st int, s string) (int, int) {
	en := st
	for ; en < len(s); en++ {
		if s[en] >= '0' && s[en] <= '9' {
			continue
		} else {
			break
		}
	}
	number, _ := strconv.Atoi(s[st : en])
	return number, en + 1
}

func decode(st int, s string, num int) (int, string) {
	var builder strings.Builder
	for i := st; i < len(s); i++ {
		if s[i] == ']' {
			tmp := builder.String()
			for j := 0; j < num-1; j++ {
				builder.WriteString(tmp)
			}
			return i, builder.String()
		} else if s[i] >= '0' && s[i] <= '9' {
			number, stringStart := findNumber(i, s)
			end, res := decode(stringStart, s, number)
			builder.WriteString(res)
			i = end
		} else {
			builder.WriteByte(s[i])
		}
	}
	return -1, ""
}
