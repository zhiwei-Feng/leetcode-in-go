package offer0065

func add(a int, b int) int {
	for b != 0 {
		c := (a & b) << 1
		a = a ^ b
		b = c
	}
	return a
}
