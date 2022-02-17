package n0046

func permute(nums []int) [][]int {
	ans := [][]int{}
	list := []int{}
	vis := make(map[int]bool)
	var dfs func(index int)
	dfs = func(index int) {
		if index == len(nums) {
			tmp := make([]int, len(list))
			copy(tmp, list)
			ans = append(ans, tmp)
			return
		}

		for _, v := range nums {
			if _, ok := vis[v]; !ok {
				vis[v] = true
				list = append(list, v)
				dfs(index + 1)
				delete(vis, v)
				list = list[:len(list)-1]
			}
		}
	}

	dfs(0)
	return ans
}
