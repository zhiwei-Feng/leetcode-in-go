package n0040

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ans := [][]int{}
	path := []int{}
	var dfs func(idx int, remain int)
	dfs = func(idx int, remain int) {
		if idx == len(candidates) || remain == 0 {
			if remain == 0 {
				ans = append(ans, append([]int{}, path...))
			}
			return
		}

		for i := idx; i < len(candidates); i++ {
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}
			if candidates[i] > remain {
				break
			}
			path = append(path, candidates[i])
			dfs(i+1, remain-candidates[i])
			path = path[:len(path)-1]
		}

	}
	dfs(0, target)

	return ans
}
