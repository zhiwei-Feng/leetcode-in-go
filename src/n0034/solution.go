package n0034

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	targetL := search(nums, target)
	if targetL < 0 || targetL >= len(nums) || nums[targetL] != target {
		return []int{-1, -1}
	}
	targetR := search(nums, target+1)
	return []int{targetL, targetR - 1}
}

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
