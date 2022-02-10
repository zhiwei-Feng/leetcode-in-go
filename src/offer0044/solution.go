package offer0044

import "strconv"

func findNthDigit(n int) int {
	if n < 10 {
		return n
	}

	digit := 1
	start := 1
	count := 9
	for n > count {
		n -= count
		start *= 10
		digit += 1
		count = 9 * start * digit
	}

	num := start + (n-1)/digit
	numStr := strconv.Itoa(num)
	res := numStr[(n-1)%digit] - '0'
	return int(res)
}
