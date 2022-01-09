package offer0030

type MinStack struct {
	Stack  []int
	TmpMin []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Stack:  make([]int, 0), // 栈主体
		TmpMin: make([]int, 0), // 存储栈中非严格递减的元素
	}
}

func (this *MinStack) Push(x int) {
	this.Stack = append(this.Stack, x)
	if len(this.TmpMin) == 0 || x <= this.TmpMin[len(this.TmpMin)-1] {
		this.TmpMin = append(this.TmpMin, x)
	}
}

func (this *MinStack) Pop() {
	popVal := this.Stack[len(this.Stack)-1]
	this.Stack = this.Stack[:len(this.Stack)-1]
	if popVal == this.TmpMin[len(this.TmpMin)-1] {
		this.TmpMin = this.TmpMin[:len(this.TmpMin)-1]
	}
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) Min() int {
	return this.TmpMin[len(this.TmpMin)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
