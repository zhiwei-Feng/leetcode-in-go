package n0394

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	ptr := 0
	stack := []string{}
	for ptr < len(s) {
		if s[ptr] >= '1' && s[ptr] <= '9' {
			// digit
			repeatNum, nextStart := getDigit(s, ptr)
			stack = append(stack, repeatNum)
			ptr = nextStart
		} else if (s[ptr] >= 'a' && s[ptr] <= 'z') || s[ptr] == '[' {
			stack = append(stack, string(s[ptr]))
			ptr++
		} else {
			// ']'
			sub := []string{}
			for stack[len(stack)-1] != "[" {
				sub = append(sub, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			repeat, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			reverseArr(sub)
			t := strings.Repeat(getString(sub), repeat)
			stack = append(stack, t)
			ptr++
		}
	}
	return getString(stack)
}

func getDigit(s string, start int) (string, int) {
	for i := start + 1; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			continue
		} else {
			return s[start:i], i
		}
	}
	return string(s[start]), start + 1
}

func getString(stack []string) string {
	var ans strings.Builder
	for _, v := range stack {
		ans.WriteString(v)
	}
	return ans.String()
}

func reverseArr(arr []string) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}
