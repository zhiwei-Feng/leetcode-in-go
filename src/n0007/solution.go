package n0007

import "math"

func reverse(x int) int {
	ans := 0
	for x != 0 {
		ans *= 10
		cur := x % 10
		ans += cur
		x /= 10
		if ans<math.MinInt32||ans>math.MaxInt32 {
			return 0
		}
	}
	return ans
}
