package offerii0011

func findMaxLength(nums []int) int {
	prevSumIdx := make(map[int]int)
	prevSumIdx[0] = -1
	prevSum := 0
	ans := 0
	for i, v := range nums {
		if v == 0 {
			prevSum += -1
		} else {
			prevSum += 1
		}
		if pi, ok := prevSumIdx[prevSum]; ok {
			ans = max(ans, i-pi)
		} else {
			prevSumIdx[prevSum] = i
		}
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
