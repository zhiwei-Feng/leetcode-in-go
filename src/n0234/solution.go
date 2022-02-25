package n0234

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	dummy := &ListNode{Next: head}
	mid, p := dummy, dummy
	for p != nil && p.Next != nil {
		p = p.Next.Next
		mid = mid.Next
	}
	qstart := head
	if p != nil {
		// 说明节点个数是偶数个, mid处于左侧的最后一个
		for dummy.Next != mid {
			nextQ := qstart.Next
			qstart.Next = dummy.Next
			dummy.Next = qstart
			qstart = nextQ
		}
		pstart := dummy.Next
		for pstart != nil && qstart != nil {
			if pstart.Val != qstart.Val {
				return false
			}
			pstart = pstart.Next
			qstart = qstart.Next
		}
		return true
	} else {
		// 奇数个，mid处于中间
		for dummy.Next != mid {
			nextQ := qstart.Next
			qstart.Next = dummy.Next
			dummy.Next = qstart
			qstart = nextQ
		}
		pstart := dummy.Next.Next
		for pstart != nil && qstart != nil {
			if pstart.Val != qstart.Val {
				return false
			}
			pstart = pstart.Next
			qstart = qstart.Next
		}
		return true
	}
}
