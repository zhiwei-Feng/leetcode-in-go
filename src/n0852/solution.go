package n0852

func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)-1
	for l < r {
		m := l + (r-l)/2
		if arr[m] > arr[m-1] && arr[m] > arr[m+1] {
			return m
		}
		if arr[m] < arr[m+1] {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
