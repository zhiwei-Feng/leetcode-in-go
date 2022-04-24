package offerii0005

func maxProduct(words []string) (ans int) {
	byteBitArr := make([]int, len(words))
	for i := range words {
		byteBit := 0
		for j := range words[i] {
			idx := words[i][j] - 'a'
			byteBit |= 1 << idx
		}
		byteBitArr[i] = byteBit
	}

	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if byteBitArr[i]&byteBitArr[j] == 0 {
				ans = max(ans, len(words[i])*len(words[j]))
			}
		}
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
