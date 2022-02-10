package offer0017

import "strings"

func printNumbers(n int) string {
	nine := 0
	start := n - 1
	num := make([]byte, n)
	loop := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	var res strings.Builder
	var dfs func(x int)
	dfs = func(x int) {
		if x == n {
			tmpS := string(num[start:])
			if tmpS != "0" {
				res.WriteString(tmpS + ",")
			}
			if n-start == nine {
				start--
			}
			return
		}

		for _, v := range loop {
			if v == '9' {
				nine++
			}
			num[x] = v
			dfs(x + 1)
		}
		nine--
	}

	dfs(0)
	l := len(res.String())
	return res.String()[:l-1]
}
