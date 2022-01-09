package offer0009

type CQueue struct {
	Size     int
	Elements []int
}

func Constructor() CQueue {
	return CQueue{
		Size:     0,
		Elements: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.Size++
	this.Elements = append(this.Elements, value)
}

func (this *CQueue) DeleteHead() int {
	if this.Size == 0 {
		return -1
	}
	this.Size--
	delHead := this.Elements[0]
	this.Elements = this.Elements[1:]
	return delHead
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
