package n0715

import "sort"

type Range struct {
	Left, Right int
}
type RangeModule struct {
	SortedRanges []Range
}

func Constructor() RangeModule {
	return RangeModule{SortedRanges: make([]Range, 0)}
}

func (this *RangeModule) AddRange(left int, right int) {
	newRange := []Range{}
	i := 0
	var mergeRangeLeft = left
	var mergeRangeRight = right
	for i < len(this.SortedRanges) {
		if this.SortedRanges[i].Right < left {
			newRange = append(newRange, this.SortedRanges[i])
			i++
		} else {
			mergeRangeLeft = min(this.SortedRanges[i].Left, mergeRangeLeft)
			break
		}
	}
	j := i
	for j < len(this.SortedRanges) {
		if this.SortedRanges[j].Left > right {
			break
		} else {
			mergeRangeRight = max(mergeRangeRight, this.SortedRanges[j].Right)
			j++
		}
	}
	newRange = append(newRange, Range{mergeRangeLeft, mergeRangeRight})
	for i := j; i < len(this.SortedRanges); i++ {
		newRange = append(newRange, this.SortedRanges[i])
	}
	this.SortedRanges = newRange
}

func (this *RangeModule) QueryRange(left int, right int) bool {
	index := sort.Search(len(this.SortedRanges), func(i int) bool {
		return this.SortedRanges[i].Left > left
	})
	if index == 0 {
		return false
	}

	shouldIncludeSet := this.SortedRanges[index-1]
	if shouldIncludeSet.Left <= left && shouldIncludeSet.Right >= right {
		return true
	} else {
		return false
	}
}

func (this *RangeModule) RemoveRange(left int, right int) {
	newRange := []Range{}
	i := 0
	for i < len(this.SortedRanges) {
		if this.SortedRanges[i].Right < left {
			newRange = append(newRange, this.SortedRanges[i])
			i++
		} else {
			break
		}
	}
	j := i
	for j < len(this.SortedRanges) {
		if this.SortedRanges[j].Left > right {
			break
		} else {
			j++
		}
	}
	for p := i; p < j; p++ {
		curRange := this.SortedRanges[p]
		if left <= curRange.Left && right >= curRange.Right {
			continue
		}
		if curRange.Left < left {
			newRange = append(newRange, Range{curRange.Left, left})
			if curRange.Right > right {
				// split two
				newRange = append(newRange, Range{right, curRange.Right})
			}
		}else {
			newRange = append(newRange, Range{right, curRange.Right})
		}
	}
	for i := j; i < len(this.SortedRanges); i++ {
		newRange = append(newRange, this.SortedRanges[i])
	}
	this.SortedRanges = newRange
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */
