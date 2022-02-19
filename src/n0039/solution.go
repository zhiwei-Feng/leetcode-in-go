package n0039

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var dfs func(target int, index int)
	path := []int{}
	ans := [][]int{}
	dfs = func(target int, index int) {
		for i := index; i < len(candidates); i++ {
			v := candidates[i]
			if target > v {
				path = append(path, v)
				dfs(target-v, i)
				path = path[:len(path)-1]
			} else if target == v {
				tmp := make([]int, len(path))
				copy(tmp, path)
				tmp = append(tmp, v)
				ans = append(ans, tmp)
				break
			} else {
				continue
			}
		}
	}

	dfs(target, 0)
	return ans
}
