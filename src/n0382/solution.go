package n0382

import "math/rand"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	List *ListNode
}

func Constructor(head *ListNode) Solution {
	sol := Solution{
		List: head,
	}
	return sol
}

func (this *Solution) GetRandom() int {
	node := this.List
	i := 1
	ans := 0
	for ; node != nil; node = node.Next {
		if rand.Intn(i) == 0 {
			ans = node.Val
		}
		i++
	}
	return ans
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
