package offer0043

func countDigitOne(n int) (ans int) {
	for k, tenk := 0, 1; n >= tenk; k++ {
		ans += (n/(tenk*10))*tenk + compute(n%(tenk*10), tenk)
		tenk *= 10
	}
	return ans
}

func compute(n int, tenk int) int {
	if n < tenk {
		return 0
	} else if n >= tenk && n < 2*tenk {
		return n - tenk + 1
	} else {
		return tenk
	}
}
