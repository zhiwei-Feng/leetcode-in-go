package n0033

func search(nums []int, target int) int {
	var (
		start = 0
		end   = len(nums) - 1
	)

	for start <= end {
		mid := start + (end-start)/2
		if target == nums[mid] {
			return mid
		}

		if nums[start] <= nums[mid] {
			//左侧有序
			if target < nums[mid] && target >= nums[start] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}

	return -1
}
