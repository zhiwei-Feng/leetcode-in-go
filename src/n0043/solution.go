package n0043

import (
	"strconv"
	"strings"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	m, n := len(num1), len(num2)
	ans := make([]int, 1000)
	for i := m-1; i >= 0; i-- {
		// num1的每一位作为基点
		base := int(num1[i] - '0')
		for j := n-1; j >= 0; j-- {
			// 遍历num2进行相乘
			mul := int(num2[j] - '0')
			curIndex := m-1-i+n-1-j
			tmp := mul*base + ans[curIndex]
			ans[curIndex] = tmp % 10
			ans[curIndex+1] += tmp / 10
		}
	}
	var ansBuilder strings.Builder
	start := 999
	for start >= 0 && ans[start] == 0 {
		start--
	}
	for i := start; i >= 0; i-- {
		ansBuilder.WriteString(strconv.Itoa(ans[i]))
	}
	return ansBuilder.String()
}
