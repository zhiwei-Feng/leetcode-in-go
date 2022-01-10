package offer0005

import "strings"

func replaceSpace(s string) string {
	var builder strings.Builder

	for _, v := range s {
		if v == ' ' {
			builder.WriteString("%20")
		} else {
			builder.WriteRune(v)
		}
	}
	return builder.String()
}
