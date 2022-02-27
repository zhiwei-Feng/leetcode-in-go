package n0008

import "math"

func myAtoi(s string) int {
	// skip space
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	if i == len(s) {
		return 0
	}
	// has flag?
	neg := false
	if s[i] == '-' || s[i] == '+' {
		neg = s[i] == '-'
		i++
	}
	ans := 0
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		cur := s[i] - '0'
		ans *= 10
		if neg {
			ans -= int(cur)
		} else {
			ans += int(cur)
		}
		if ans < math.MinInt32 {
			return math.MinInt32
		} else if ans > math.MaxInt32 {
			return math.MaxInt32
		}
		i++
	}
	return ans
}
