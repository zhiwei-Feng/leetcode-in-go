package n0029

import (
	"math"
)

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend == math.MinInt32 {
			return math.MaxInt32
		} else {
			return -dividend
		}
	}

	a, b := dividend, divisor
	sign := true
	if (a > 0 && b < 0) || (a < 0 && b > 0) {
		sign = false
	}
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	res := div(a, b)
	if sign {
		if res > math.MaxInt32 {
			res = math.MaxInt32
		}
		return res
	}
	return -res

}

// a/b, a<0, b<0
func div(a, b int) int {
	if a < b {
		return 0
	}
	count := 1
	result := b
	for result+result <= a {
		result += result
		count += count
	}
	return count + div(a-result, b)
}
