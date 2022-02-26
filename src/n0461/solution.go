package n0461

func hammingDistance(x int, y int) int {
	t := x ^ y
	dis := 0
	for t != 0 {
		dis++
		t = t & (t - 1)
	}
	return dis
}
