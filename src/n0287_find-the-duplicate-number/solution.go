package n0287findtheduplicatenumber

func findDuplicate_binary_search(nums []int) int {
	low, high := 1, len(nums)
	for low < high {
		mid := low + (high-low)/2
		cnt := 0
		for _, x := range nums {
			if x <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

func findDuplicate(nums []int) int {
	var (
		slow, fast = 0, 0
	)
	slow = nums[slow]
	fast = nums[nums[fast]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
