package n1231

import "math"

func maximizeSweetness(sweetness []int, k int) int {
	sum := 0
	minV := math.MaxInt16

	for _, v := range sweetness {
		sum += v
		if v < minV {
			minV = v
		}
	}

	low, high := minV, sum/(k+1)+1
	ans := 0
	for low < high {
		mid := low + (high-low)/2
		if countPart(sweetness, mid) < k+1 {
			high = mid
		} else {
			ans = mid
			low = mid + 1
		}
	}
	return ans
}

// how many part can be divide when sum of each part must greater than sweety
// @param sweety is minimum sum of each part
func countPart(sweetness []int, sweety int) int {
	sum := 0
	count := 0
	for _, v := range sweetness {
		sum += v
		if sum >= sweety {
			count++
			sum = 0
		}
	}
	return count
}
