package n0115

func numDistinct(s string, t string) int {
	path := []byte{}
	ans := 0
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == len(s) || len(path) == len(t) {
			if string(path) == t {
				ans++
			}
			return
		}

		// 不选择当前的
		dfs(idx + 1)
		// 选择当前的
		path = append(path, s[idx])
		dfs(idx + 1)
		path = path[:len(path)-1]
	}
	dfs(0)
	return ans
}
