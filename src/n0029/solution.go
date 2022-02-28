package n0029

import "math"

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return min(dividend, math.MaxInt32)
	}
	if dividend < 0 {
		return -divide(-dividend, divisor)
	}
	if divisor < 0 {
		return -divide(dividend, -divisor)
	}
	if dividend < divisor {
		return 0
	}

	ans := 1
	d := divisor
	for (d + d) <= dividend {
		d = d + d
		ans = ans + ans
	}
	for d <= dividend {
		d += divisor
		ans++
	}
	if ans == 1 {
		return 1
	}
	return min(ans-1, math.MaxInt32)
}

// x*y>=z是否成立
func quickAdd(x, y, z int) bool {
	result := 0
	if y&1 == 1 {
		result += x
		y--
	}
	// y为偶数
	
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// func divide(dividend int, divisor int) int {
// 	if dividend == 0 {
// 		return 0
// 	}
// 	if divisor == 1 {
// 		return min(dividend, math.MaxInt32)
// 	}
// 	if dividend < 0 {
// 		return -divide(-dividend, divisor)
// 	}
// 	if divisor < 0 {
// 		return -divide(dividend, -divisor)
// 	}
// 	if dividend < divisor {
// 		return 0
// 	}

// 	ans := 1
// 	d := divisor
// 	for (d + d) <= dividend {
// 		d = d + d
// 		ans = ans + ans
// 	}
// 	for d <= dividend {
// 		d += divisor
// 		ans++
// 	}
//     if ans == 1 {
// 		return 1
// 	}
// 	return min(ans-1, math.MaxInt32)
// }

// func min(i, j int) int {
// 	if i < j {
// 		return i
// 	}
// 	return j
// }
