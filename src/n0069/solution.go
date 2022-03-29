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

func mySqrtFloat(x float64) float64 {
	if x <= 1 {
		return x
	}
	const delta = 1e-10
	var l, r float64 = 1, x
	for l < r {
		mid := (l + r) / 2
		if abs(mid*mid-x)<delta*delta{
			return mid
		}else if mid*mid<x {
			l = mid+delta
		}else {
			r = mid-delta
		}
	}
	return l
}

func abs(x float64) float64{
	if x < 0 {
		return -x
	}
	return x
}
