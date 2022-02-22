package n0155

type Node struct {
	Val  int
	Min  int
	Next *Node
}

type MinStack struct {
	StackHead *Node
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if this.StackHead == nil {
		this.StackHead = &Node{Val: val, Min: val}
	} else {
		this.StackHead = &Node{Val: val, Min: min(val, this.StackHead.Min), Next: this.StackHead}
	}
}

func (this *MinStack) Pop() {
	this.StackHead = this.StackHead.Next
}

func (this *MinStack) Top() int {
	return this.StackHead.Val
}

func (this *MinStack) GetMin() int {
	return this.StackHead.Min
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// type MinStack struct {
// 	Stack []int
// 	MinQ  []int // 非严格递减
// }

// func Constructor() MinStack {
// 	return MinStack{
// 		Stack: make([]int, 0),
// 		MinQ:  make([]int, 0),
// 	}
// }

// func (this *MinStack) Push(val int) {
// 	this.Stack = append(this.Stack, val)
// 	if len(this.MinQ) == 0 || val <= this.MinQ[len(this.MinQ)-1] {
// 		this.MinQ = append(this.MinQ, val)
// 	}
// }

// func (this *MinStack) Pop() {
// 	n := len(this.Stack)
// 	x := this.Stack[n-1]
// 	this.Stack = this.Stack[:n-1]
// 	if x == this.MinQ[len(this.MinQ)-1] {
// 		this.MinQ = this.MinQ[:len(this.MinQ)-1]
// 	}
// }

// func (this *MinStack) Top() int {
// 	return this.Stack[len(this.Stack)-1]
// }

// func (this *MinStack) GetMin() int {
// 	return this.MinQ[len(this.MinQ)-1]
// }
