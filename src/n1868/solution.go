package n1868

// 转化区间交集问题的解法
func findRLEArray(encoded1 [][]int, encoded2 [][]int) [][]int {
	prodEncoded := [][]int{}
	st1, st2 := 0, 0
	i, j := 0, 0
	for i < len(encoded1) && j < len(encoded2) {
		overlapX := max(st1, st2)
		overlapY := min(st1+encoded1[i][1]-1, st2+encoded2[j][1]-1)

		if overlapX <= overlapY {
			prod := encoded1[i][0] * encoded2[j][0]
			length := overlapY - overlapX + 1
			if len(prodEncoded) == 0 || prod != prodEncoded[len(prodEncoded)-1][0] {
				prodEncoded = append(prodEncoded, []int{prod, length})
			} else {
				prodEncoded[len(prodEncoded)-1][1] += length
			}

		}

		if st1+encoded1[i][1] < st2+encoded2[j][1] {
			st1 += encoded1[i][1]
			i++
		} else {
			st2 += encoded2[j][1]
			j++
		}
	}

	return prodEncoded
}

func min(i1, i2 int) int {
	if i1 < i2 {
		return i1
	}
	return i2
}

func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}
