package n2047

import "strings"

func countValidWords(sentence string) int {
	tokens := strings.Split(sentence, " ")
	ans := 0
	for _, token := range tokens {
		if token == "" {
			continue
		}

		flag := true
		andCount := 0
		for i := 0; i < len(token); i++ {
			if !validByte(token[i]) {
				flag = false
				break
			}
			if token[i] == '-' {
				if i == 0 || i == len(token)-1 {
					flag = false
					break
				}
				if token[i-1] < 'a' || token[i-1] > 'z' || token[i+1] < 'a' || token[i+1] > 'z' {
					flag = false
					break
				}
				andCount++
			}

			if (token[i] == '!' || token[i] == '.' || token[i] == ',') && i != len(token)-1 {
				flag = false
				break
			}

		}
		if flag && andCount < 2 {
			ans++
		}
	}

	return ans
}

func validByte(x byte) bool {
	var (
		islowerByte bool = x >= 'a' && x <= 'z'
		isAnd       bool = x == '-'
		isEndToken  bool = x == '!' || x == '.' || x == ','
	)
	return islowerByte || isAnd || isEndToken
}
