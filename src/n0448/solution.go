package n0448

func findDisappearedNumbers(nums []int) []int {
	ans := []int{}
	n := len(nums)
	for _, v := range nums {
		index := (v - 1) % n
		nums[index] += n
	}
	for i, v := range nums {
		if v <= n {
			ans = append(ans, i+1)
		}
	}
	return ans
}
