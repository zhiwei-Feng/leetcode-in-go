package offer0053_1

func search(nums []int, target int) int {
	low, high := 0, len(nums)
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] >= target {
			high = mid
		} else {
			low = mid + 1
		}
	}
	start := low
	high = len(nums)
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] <= target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	end := low
	return end - start
}
