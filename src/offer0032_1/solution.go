package offer0032_1

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := []int{}
	q := []*TreeNode{}
	q = append(q, root)
	for len(q) != 0 {
		top := q[0]
		q = q[1:]
		ans = append(ans, top.Val)
		if top.Left != nil {
			q = append(q, top.Left)
		}
		if top.Right != nil {
			q = append(q, top.Right)
		}
	}

	return ans
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}
