package main

import (
	"fmt"
)

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
	p := head
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

func main() {
	head := &Node{}
	tail := head
	for i := 1; i <= 8; i++ {
		tail.Next = &Node{Val: i}
		tail = tail.Next
	}
	for p := head.Next; p != nil; p = p.Next {
		fmt.Print(p.Val)
	}
	fmt.Println()
	newHead := Solver(head.Next, 3)
	for p := newHead; p != nil; p = p.Next {
		fmt.Print(p.Val)
	}
	fmt.Println()
}
