package offer0055_1

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	q := []*TreeNode{root}
	ans := 0
	for len(q) != 0 {
		ans++
		levelLen := len(q)
		for i := 0; i < levelLen; i++ {
			cur := q[0]
			q = q[1:]
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
	}
	return ans
}
