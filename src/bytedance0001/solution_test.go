package bytedance0001

import (
	"fmt"
	"testing"
)

func TestSolver(t *testing.T) {
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
	newHead := Solver(head.Next, 5)
	for p := newHead; p != nil; p = p.Next {
		fmt.Print(p.Val)
	}
	fmt.Println()
}
