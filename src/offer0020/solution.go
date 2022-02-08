package offer0020

import "strings"

func isNumber(s string) bool {
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return false
	}
	hasDot := false
	hasInt := false
	for i := 0; i < len(s); i++ {
		if s[i] == '+' || s[i] == '-' {
			if i == 0 {
				continue
			} else {
				return false
			}
		}
		if s[i] == '.' {
			if hasDot {
				return false
			}

			if !hasInt && (i+1 >= len(s) || (s[i+1] < '0' || s[i+1] > '9')) {
				return false
			}
			hasDot = true
		} else if s[i] == 'e' || s[i] == 'E' {
			if !hasInt {
				return false
			}
			return isInteger(s[i+1:])
		} else if s[i] >= '0' && s[i] <= '9' {
			hasInt = true
			continue
		} else {
			return false
		}
	}
	return true
}

func isInteger(s string) bool {
	if len(s) == 0 {
		return false
	}
	hasInt := false
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			if i == 0 && (s[i] == '+' || s[i] == '-') {
				continue
			}
			return false
		} else {
			hasInt = true
		}
	}

	return hasInt
}
