package bytedance0001

/*
这其实是一道变形的链表反转题，大致描述如下 给定一个单链表的头节点 head,实现一个调整单链表的函数，使得每K个节点之间为一组进行逆序，并且从链表的尾部开始组起，头部剩余节点数量不够一组的不需要逆序。（不能使用队列或者栈作为辅助）

例如：

链表:1->2->3->4->5->6->7->8->null, K = 3。那么 6->7->8，3->4->5，1->2各位一组。调整后：1->2->5->4->3->8->7->6->null。其中 1，2不调整，因为不够一组。
*/
type Node struct {
	Val  int
	Next *Node
}

func Solver(head *Node, k int) *Node {
	// reverse
	newHead := reverse(head)
	dummy := &Node{}
	prevH := dummy
	H := newHead
	p := newHead
	count := 0
	for p != nil {
		count++
		if count == k {
			nextP := p.Next
			p.Next = nil // 截断
			newH := reverse(H)
			prevH.Next = newH
			H.Next = nextP
			p = nextP
			count = 0
			prevH = H
			H = H.Next
		} else {
			p = p.Next
		}
	}
	finalHead := reverse(dummy.Next)
	return finalHead
}

func reverse(head *Node) *Node {
	p := &Node{Next: nil}
	q := head
	for q != nil {
		nextq := q.Next
		q.Next = p.Next
		p.Next = q
		q = nextq
	}
	return p.Next
}
