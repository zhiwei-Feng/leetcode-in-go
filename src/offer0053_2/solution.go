package offer0053_2

// 位运算和二分查找都可
func missingNumber(nums []int) int {
	low, high := 0, len(nums)
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] == mid {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}
