package n0390

func lastRemaining(n int) int {
	cnt := n
	step := 1
	start := 1
	left := true
	for cnt != 1 {
		if cnt%2 != 0 {
			start += step
		} else {
			if left {
				start += step
			}
		}
		step *= 2
		cnt /= 2
		left = !left
	}
	return start
}
