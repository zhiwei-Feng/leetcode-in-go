package n0069

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	low, high := 1, x
	for low < high {
		mid := low + (high-low+1)/2
		if mid > x/mid {
			high = mid - 1
		} else {
			low = mid
		}
	}
	return low
}
