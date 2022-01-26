package n2013

type DetectSquares struct {
	cnt map[int]map[int]int
}

func Constructor() DetectSquares {
	return DetectSquares{make(map[int]map[int]int)}
}

func (this *DetectSquares) Add(point []int) {
	x, y := point[0], point[1]
	if this.cnt[y] == nil {
		this.cnt[y] = map[int]int{}
	}
	this.cnt[y][x]++
}

func (this *DetectSquares) Count(point []int) int {
	ans := 0
	x, y := point[0], point[1]
	if this.cnt[y] == nil {
		return 0
	}
	yCnt := this.cnt[y]
	for col, colCnt := range this.cnt {
		if col != y {
			d := col - y
			ans += colCnt[x] * yCnt[x+d] * colCnt[x+d]
			ans += colCnt[x] * yCnt[x-d] * colCnt[x-d]
		}
	}
	return ans
}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */
