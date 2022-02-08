package offer0067

import (
	"math"
	"strings"
)

func strToInt(str string) int {
	str = strings.Trim(str, " ")
	if len(str) == 0 {
		return 0
	}

	negative := false
	hasValidNum := false
	ans := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '+' || str[i] == '-' {
			if i == 0 {
				negative = str[i] == '-'
				continue
			} else if hasValidNum {
				break
			} else {
				return 0
			}
		}

		if validNum(str[i]) {
			hasValidNum = true
			ans *= 10
			ans += int(str[i] - '0')
			if ans > math.MaxInt32+1 {
				break
			} 
			continue
		} else if hasValidNum {
			break
		} else {
			return 0
		}
	}

	if negative {
		ans = -ans
	}

	if ans > math.MaxInt32 {
		ans = math.MaxInt32
	} else if ans < math.MinInt32 {
		ans = math.MinInt32
	}
	return ans
}

func validNum(b byte) bool {
	return b >= '0' && b <= '9'
}
