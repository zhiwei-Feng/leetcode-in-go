package n0141

type ListNode struct {
	Val int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
    fast, slow := head, head
    for fast!=nil&&fast.Next!=nil {
        fast = fast.Next.Next
        slow = slow.Next
        if fast==slow{
            break
        }
    }
    return fast!=nil && fast.Next!=nil
}