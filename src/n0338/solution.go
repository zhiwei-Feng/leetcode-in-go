package n0338

func countBits(n int) []int {
	hightbit := 0
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			hightbit = i
		}
		bits[i] = bits[i-hightbit] + 1
	}
	return bits
}
