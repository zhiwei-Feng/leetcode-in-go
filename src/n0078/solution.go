package n0078

func subsets(nums []int) [][]int {
	ans := [][]int{{}}
	for _, v := range nums {
		for _, win := range ans {
			tmp := make([]int, len(win))
			copy(tmp, win)
			tmp = append(tmp, v)
			ans = append(ans, tmp)
		}
	}
	return ans
}
