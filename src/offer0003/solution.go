package offer0003

func findRepeatNumber(nums []int) int {
	count := make([]int, len(nums))
	for _, v := range nums {
		count[v]++
		if count[v] > 1 {
			return v
		}
	}
	return -1
}
