package interview0019

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	match := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	f := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2]
				if match(i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if match(i, j) {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}

// func isMatch(s string, p string) bool {
// 	if len(p) == 0 {
// 		return len(s) == 0
// 	}

// 	firstMatch := len(s) != 0 && (s[0] == p[0] || p[0] == '.')
// 	if len(p) >= 2 && p[1] == '*' {
// 		return (firstMatch && isMatch(s[1:], p)) || isMatch(s, p[2:])
// 	} else {
// 		return firstMatch && isMatch(s[1:], p[1:])
// 	}
// }
