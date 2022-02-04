package offer0016

// 快速幂
func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	res := 1.0
	if n < 0 {
		x = 1 / x
		n = -n
	}
	for n > 0 {
		if (n & 1) == 1 {
			res *= x
		}
		x *= x
		n = n >> 1
	}
	return res
}

// func myPow(x float64, n int) float64 {
// 	if n == 0 {
// 		return 1
// 	}
// 	if n < 0 {
// 		return 1 / myPow(x, -n)
// 	}
// 	if n%2 == 0 {
// 		ans := myPow(x, n/2)
// 		return ans * ans
// 	} else {
// 		return x * myPow(x, n-1)
// 	}
// }
