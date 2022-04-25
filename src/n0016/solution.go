package n0016

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	closest := math.MaxInt64 / 2
	sort.Ints(nums)
	for a := 0; a < len(nums)-2; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}

		c := len(nums) - 1
		for b := a + 1; b < c; b++ {
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}
			for b < c && nums[a]+nums[b]+nums[c] > target {
				if isMoreClosed(nums[a]+nums[b]+nums[c], closest, target) {
					closest = nums[a] + nums[b] + nums[c]
				}
				c--
			}
			if b == c {
				break
			}
			if isMoreClosed(nums[a]+nums[b]+nums[c], closest, target) {
				closest = nums[a] + nums[b] + nums[c]
			}
		}
	}
	return closest
}

func isMoreClosed(a, b int, target int) bool {
	// return true if a is more closed to target than b
	return abs(a-target) < abs(b-target)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
