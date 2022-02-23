package n0301

func removeInvalidParentheses(s string) []string {
	ll, rr := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			ll++
		} else if s[i] == ')' {
			if ll > 0 {
				ll--
			} else {
				rr++
			}
		}
	}
	ans := []string{}
	helper(&ans, s, 0, ll, rr)
	return ans
}

func isValid(str string) bool {
	cnt := 0
	for _, ch := range str {
		if ch == '(' {
			cnt++
		} else if ch == ')' {
			cnt--
			if cnt < 0 {
				return false
			}
		}
	}
	return cnt == 0
}

func helper(ans *[]string, str string, start, ll, rr int) {
	if ll == 0 && rr == 0 {
		if isValid(str) {
			*ans = append(*ans, str)
		}
		return
	}

	for i := start; i < len(str); i++ {
		if i != start && str[i] == str[i-1] {
			continue
		}
		if ll+rr > len(str)-i { // 后续字符无法满足减去ll+rr个字符
			return
		}

		// try discard left
		if ll > 0 && str[i] == '(' {
			helper(ans, str[:i]+str[i+1:], i, ll-1, rr)
		}
		if rr > 0 && str[i] == ')' {
			helper(ans, str[:i]+str[i+1:], i, ll, rr-1)
		}
	}
}

func removeInvalidParentheses_(s string) []string {
	var (
		maxScore int //最终字符串每有一个'('，则+1， 一个')'则-1
		maxLen   int
		ansMap   = make(map[string]int) // 处理重复答案
		ans      = []string{}
		dfs      func(curP int, tmpS []rune, delL, delR int, score int)
	)

	// [1]: 预处理s得到合法字符串的最大长度
	ll, rr := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			ll++
		} else if s[i] == ')' {
			if ll > 0 {
				ll--
			} else {
				rr++
			}
		}
	}
	maxLen = len(s) - ll - rr

	// [2]: 得到score的最大值
	l, r := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			l++
		} else if s[i] == ')' {
			r++
		}
	}
	maxScore = min(l, r)

	// [3]: dfs
	// delL, delR为当前还需要删除几个左括号和右括号
	dfs = func(curP int, tmpS []rune, delL, delR, score int) {
		if delL < 0 || delR < 0 || score < 0 || score > maxScore {
			return
		}
		if delL == 0 && delR == 0 && len(tmpS) == maxLen {
			ansMap[string(tmpS)]++
		}
		if curP == len(s) {
			return
		}

		curRune := rune(s[curP])
		if curRune == '(' {
			dfs(curP+1, append(tmpS, '('), delL, delR, score+1)
			dfs(curP+1, tmpS, delL-1, delR, score)
		} else if curRune == ')' {
			dfs(curP+1, append(tmpS, ')'), delL, delR, score-1)
			dfs(curP+1, tmpS, delL, delR-1, score)
		} else {
			dfs(curP+1, append(tmpS, curRune), delL, delR, score)
		}
	}
	dfs(0, []rune{}, ll, rr, 0)
	for k := range ansMap {
		ans = append(ans, k)
	}
	return ans
}

func removeInvalidParentheses_bruteDFS(s string) []string {
	var (
		maxScore int
		maxLen   int
		ansMap   = make(map[string]int)
		ans      = []string{}
		dfs      func(curP int, tmpS []rune, score int)
	)

	l, r := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			l++
		} else if s[i] == ')' {
			r++
		}
	}
	maxScore = min(l, r)

	dfs = func(curP int, tmpS []rune, score int) {
		if score < 0 || score > maxScore {
			return
		}
		if curP == len(s) {
			if score == 0 && len(tmpS) >= maxLen {
				if len(tmpS) > maxLen {
					ansMap = make(map[string]int)
				}
				ansMap[string(tmpS)]++
				maxLen = len(tmpS)
			}
			return
		}

		curRune := rune(s[curP])
		if curRune == '(' {
			dfs(curP+1, append(tmpS, '('), score+1)
			dfs(curP+1, tmpS, score)
		} else if curRune == ')' {
			dfs(curP+1, append(tmpS, ')'), score-1)
			dfs(curP+1, tmpS, score)
		} else {
			dfs(curP+1, append(tmpS, curRune), score)
		}
	}
	dfs(0, []rune{}, 0)

	for k := range ansMap {
		ans = append(ans, k)
	}

	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
