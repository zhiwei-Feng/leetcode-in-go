package n0201

import "math"

func rangeBitwiseAnd(left int, right int) int {
	// Brian Kernighan
	for left < right {
		right &= right - 1
	}

	return right
}

func rangeBitwiseAnd_rightmove(left int, right int) int {
	shift := 0
	for left < right {
		left = left >> 1
		right = right >> 1
		shift += 1
	}

	return left << shift
}

func rangeBitwiseAnd_brute(left int, right int) int {
	ans := math.MaxInt32
	for i := left; i <= right; i++ {
		ans = ans & i
	}
	return ans
}
