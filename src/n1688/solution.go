package n1688

func numberOfMatches(n int) int {
	memory := make(map[int]int)
	return match(n, memory)
}

func match(n int, mem map[int]int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return n - 1
	}
	if v, ok := mem[n]; ok {
		return v
	}

	var res int
	if n%2 == 0 {
		res = n/2 + match(n/2, mem)
	} else {
		res = (n-1)/2 + match((n-1)/2+1, mem)
	}
	mem[n] = res
	return res
}
