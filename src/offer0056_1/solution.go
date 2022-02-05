package offer0056_1

func singleNumbers(nums []int) []int {
	c := 0
	for _, v := range nums {
		c ^= v
	}

	judge := c & -c // 获取二进制最右侧1
	a, b := 0, 0
	for _, v := range nums {
		if (v & judge) == judge {
			a ^= v
		} else {
			b ^= v
		}
	}
	return []int{a, b}
}
