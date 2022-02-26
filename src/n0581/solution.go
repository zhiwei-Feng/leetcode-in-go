package n0581

import "math"

func findUnsortedSubarray(nums []int) int {
	var (
		maxL  = math.MinInt64
		minR  = math.MaxInt64
		left  = -1
		right = -1
	)
	for i := 0; i < len(nums); i++ {
		if nums[i]<maxL {
			right = i
		}else {
			maxL = nums[i]
		}
		if nums[len(nums)-i-1]>minR {
			left = len(nums)-i-1
		}else{
			minR = nums[len(nums)-i-1]
		}
	}
	if right == -1 {
		return 0
	}
	return right - left + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
