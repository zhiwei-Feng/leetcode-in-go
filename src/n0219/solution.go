package n0219

func containsNearbyDuplicate(nums []int, k int) bool {
	recent := make(map[int]int)
	for j := 0; j < len(nums); j++ {
		if i, ok := recent[nums[j]]; !ok {
			recent[nums[j]] = j
			continue
		} else if j-i <= k {
			return true
		}
		recent[nums[j]] = j
	}
	return false
}
