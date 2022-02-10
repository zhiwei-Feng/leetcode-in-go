package offer0014_2

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	mod := 1000000007
	a, b := n/3, n%3
	if b == 0 {
		return remainder(3, a, mod)
	} else if b == 1 {
		return (remainder(3, a-1, mod) * 4) % mod
	} else {
		return (remainder(3, a, mod) * 2) % mod
	}
}

// x^a % p
func remainder(x, a, p int) int {
	rem := 1
	for a > 0 {
		if a%2 == 1 {
			rem = (rem * x) % p
		}
		x = (x * x) % p
		a /= 2
	}
	return rem
}
