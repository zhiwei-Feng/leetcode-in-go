package n0038

import (
	"fmt"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	prev := "1"
	for i := 2; i <= n; i++ {
		var builder strings.Builder
		count := 0
		for j := range prev {
			if j > 0 && prev[j] != prev[j-1] {
				builder.WriteString(fmt.Sprintf("%d%c", count, prev[j-1]))
				count=1
			} else {
				count++
			}
		}
		if count > 0 {
			builder.WriteString(fmt.Sprintf("%d%d", count, prev[len(prev)-1]-'0'))
		}
		prev = builder.String()
	}
	return prev
}
