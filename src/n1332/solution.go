package n1332

func removePalindromeSub(s string) int {
	// 仅有ab，则最多两次删除即可删完
	// 全部删除a，再全部删除b
	if isPalindrome(s) {
		return 1
	}
	return 2
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}
