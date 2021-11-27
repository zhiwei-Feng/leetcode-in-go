package n5924

func minCost(startPos []int, homePos []int, rowCosts []int, colCosts []int) int {
	if startPos[0] == homePos[0] && startPos[1] == homePos[1] {
		return 0
	}
	cost := 0
	if startPos[0] <= homePos[0] {
		for i := startPos[0]; i < homePos[0]; i++ {
			cost += rowCosts[i+1]
		}
	} else {
		for i := startPos[0]; i > homePos[0]; i-- {
			cost += rowCosts[i-1]
		}
	}

	if startPos[1] <= homePos[1] {
		for j := startPos[1]; j < homePos[1]; j++ {
			cost += colCosts[j+1]
		}
	} else {
		for j := startPos[1]; j > homePos[1]; j-- {
			cost += colCosts[j-1]
		}
	}

	return cost
}
