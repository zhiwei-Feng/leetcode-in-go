package n0136

func singleNumber(nums []int) int {
	ans := 0
	for _, v := range nums {
		ans ^= v
	}
	return ans
}
