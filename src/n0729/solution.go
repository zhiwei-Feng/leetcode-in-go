package n0729

type MyCalendar struct {
	Root *Interval
}

type Interval struct {
	Start, End  int
	Left, Right *Interval
}

func (cur *Interval) insert(newInterval *Interval) bool {
	if newInterval.Start >= cur.End {
		if cur.Right == nil {
			cur.Right = newInterval
			return true
		}
		return cur.Right.insert(newInterval)
	} else if newInterval.End <= cur.Start {
		if cur.Left == nil {
			cur.Left = newInterval
			return true
		}
		return cur.Left.insert(newInterval)
	} else {
		return false
	}
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (this *MyCalendar) Book(start int, end int) bool {
	if this.Root == nil {
		this.Root = &Interval{start, end, nil, nil}
		return true
	}
	return this.Root.insert(&Interval{start, end, nil, nil})
}
