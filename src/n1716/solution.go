package n1716

func totalMoney(n int) int {
	total := 0
	day := 1
	lastMon := 0
	value := lastMon

	for day <= n {
		if day%7 == 1 {
			value = lastMon + 1
			lastMon = value
		}
		total += value
		value++
		day++
	}

	return total
}
