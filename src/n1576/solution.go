package n1576

import "strings"

func modifyString(s string) string {
	var ans strings.Builder
	var last byte = '_'
	for i := 0; i < len(s); i++ {
		if s[i] != '?' {
			ans.WriteByte(s[i])
			last = s[i]
			continue
		}

		for v := 'a'; v <= 'z'; v++ {
			if last != 'z' && v == rune(last) {
				continue
			}
			if i+1 < len(s) && rune(s[i+1]) == v {
				continue
			}
			ans.WriteRune(v)
			last = byte(v)
			break
		}
	}

	return ans.String()
}
