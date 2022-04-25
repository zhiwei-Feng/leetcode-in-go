package offerii0007

import "sort"

func threeSum(nums []int) (ans [][]int) {
	sort.Ints(nums)
	for a := 0; a < len(nums)-2; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		target := -nums[a]
		c := len(nums) - 1
		for b := a + 1; b < c; b++ {
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}
			for b < c && nums[b]+nums[c] > target {
				c--
			}
			if b == c {
				break
			}
			if nums[b]+nums[c] == target {
				ans = append(ans, []int{nums[a], nums[b], nums[c]})
			}
		}
	}
	return
}
