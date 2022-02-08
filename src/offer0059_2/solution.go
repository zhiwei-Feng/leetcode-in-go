package offer0059_2

type MaxQueue struct {
	queue []int // 原始队列
	deque []int // 非严格递减队列
}

func Constructor() MaxQueue {
	return MaxQueue{make([]int, 0), make([]int, 0)}
}

func (this *MaxQueue) Max_value() int {
	if len(this.queue) == 0 {
		return -1
	}
	return this.deque[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.queue = append(this.queue, value)
	for len(this.deque) > 0 && this.deque[len(this.deque)-1] < value {
		this.deque = this.deque[:len(this.deque)-1]
	}
	this.deque = append(this.deque, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}
	x := this.queue[0]
	this.queue = this.queue[1:]
	if x == this.deque[0] {
		this.deque = this.deque[1:]
	}
	return x
}
