package bytedance0005

func sqrt(x int) float64 {
	var l, r float64 = 0, float64(x)
	delta := 0.000001*0.000001
	for l < r {
		mid := l + (r-l)/2
		if abs((mid*mid)-float64(x)) <= delta {
			return mid
		} else if mid*mid > float64(x) {
			r = mid
		} else {
			l = mid
		}
	}
	return l
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
