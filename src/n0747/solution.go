package n0747

func dominantIndex(nums []int) int {
	total := 0
	maxV := -1
	secMaxV := -1
	maxVInd := -1

	for i, v := range nums {
		total += v
		if v > maxV {
			secMaxV = maxV
			maxV = v
			maxVInd = i
		} else if v > secMaxV {
			secMaxV = v
		}
	}

	if maxV >= 2*secMaxV {
		return maxVInd
	}
	return -1
}
