package n0560

func subarraySum(nums []int, k int) int {
	count := 0
	mpre := make(map[int]int)
	mpre[0] = 1 // 边界
	pre := 0    // prev[i-1]
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if v, ok := mpre[pre-k]; ok {
			count += v
		}
		mpre[pre]++
	}
	return count
}
