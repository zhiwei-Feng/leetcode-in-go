package n0260

func singleNumber(nums []int) []int {
	xor := nums[0]
	for i := 1; i < len(nums); i++ {
		xor ^= nums[i]
	}
	mostRightOne := xor & (-xor)
	// 通过mostRightOne将nums划分为两个部分
	n1, n2 := 2<<16-1, 2<<16-1
	for _, v := range nums {
		if mostRightOne&v == 0 {
			n1 = n1 ^ v
		} else {
			n2 = n2 ^ v
		}
	}

	return []int{n1, n2}
}
