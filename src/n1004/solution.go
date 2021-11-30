package n1004

func longestOnes(nums []int, k int) int {
	winSt := 0
	winCostK := 0
	ans := 0
	for winEn := 0; winEn < len(nums); winEn++ {
		if nums[winEn] == 0 {
			winCostK++
		}
		for winCostK > k {
			if nums[winSt] == 0 {
				winCostK--
			}
			winSt++
		}
		if winEn-winSt+1 > ans {
			ans = winEn - winSt + 1
		}
	}

	return ans
}
