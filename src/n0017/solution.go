package n0017

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	mapping := [][]byte{
		{},
		{},
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
		{'j', 'k', 'l'},
		{'m', 'n', 'o'},
		{'p', 'q', 'r', 's'},
		{'t', 'u', 'v'},
		{'w', 'x', 'y', 'z'},
	}

	ans := []string{}
	path := []byte{}
	var dfs func(index int)
	dfs = func(index int) {
		if index == len(digits) {
			tmp := make([]byte, len(path))
			copy(tmp, path)
			ans = append(ans, string(tmp))
			return
		}

		curNum := digits[index] - '0'
		for _, v := range mapping[curNum] {
			path = append(path, v)
			dfs(index + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}
