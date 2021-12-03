package n0752

func openLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}
	deadMap := make(map[string]int)
	for _, ds := range deadends {
		deadMap[ds]++
	}
	if _, ok := deadMap["0000"]; ok {
		return -1
	}
	seenMap := make(map[string]int)
	type Pair struct {
		Value []byte
		Step  int
	}
	queue := []Pair{}
	queue = append(queue, Pair{[]byte{'0', '0', '0', '0'}, 0})
	seenMap["0000"]++
	for len(queue) > 0 {
		cur, step := queue[0].Value, queue[0].Step
		queue = queue[1:]
		for _, cand := range getCandidates(cur) {
			if string(cand) == target {
				return step + 1
			}
			_, isDead := deadMap[string(cand)]
			_, isSeen := seenMap[string(cand)]
			if !isDead && !isSeen {
				queue = append(queue, Pair{cand, step + 1})
				seenMap[string(cand)]++
			}
		}
	}
	return -1
}

func getCandidates(current []byte) [][]byte {
	ans := [][]byte{}
	for i := 0; i < 4; i++ {
		nextTmp := make([]byte, len(current))
		prevTmp := make([]byte, len(current))
		copy(nextTmp, current)
		copy(prevTmp, current)
		nextTmp[i] = nextNum(nextTmp[i])
		ans = append(ans, nextTmp)
		prevTmp[i] = prevNum(prevTmp[i])
		ans = append(ans, prevTmp)
	}
	return ans
}

func nextNum(num byte) byte {
	if num == '9' {
		return '0'
	} else {
		return num + 1
	}
}
func prevNum(num byte) byte {
	if num == '0' {
		return '9'
	} else {
		return num - 1
	}
}
