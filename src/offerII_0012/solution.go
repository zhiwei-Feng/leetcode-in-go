package offerii0012

func pivotIndex(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	prevSum := 0
	for i, v := range nums {
		if prevSum == total-prevSum-v {
			return i
		}
		prevSum += v
	}
	return -1
}
