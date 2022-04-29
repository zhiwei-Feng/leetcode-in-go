package offerii0010

func subarraySum(nums []int, k int) int {
	prevCount := make(map[int]int)
	prevCount[0] = 1
	prevSum := 0
	ans := 0
	for _, v := range nums {
		prevSum += v
		ans += prevCount[prevSum-k]
		prevCount[prevSum]++
	}
	return ans
}
