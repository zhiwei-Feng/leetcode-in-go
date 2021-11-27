package n5923

func minimumBuckets(street string) int {
	ans := 0
	last := -2 // to avoid the index 0
	for i := 0; i < len(street); i++ {
		if street[i] == 'H' {
			if last == i-1 {
				continue
			}
			if i+1 < len(street) && street[i+1] == '.' {
				last = i + 1
				ans++
			} else if i >= 1 && street[i-1] == '.' {
				last = i - 1
				ans++
			} else {
				return -1
			}
		}
	}

	return ans
}
