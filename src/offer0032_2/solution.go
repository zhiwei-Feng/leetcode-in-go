package offer0032_2

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := [][]int{}
	q := []*TreeNode{}
	q = append(q, root)
	for len(q) != 0 {
		levelN := len(q)
		tmp := []int{}
		for i := 0; i < levelN; i++ {
			tmp = append(tmp, q[0].Val)
			if q[0].Left != nil {
				q = append(q, q[0].Left)
			}
			if q[0].Right != nil {
				q = append(q, q[0].Right)
			}
			q = q[1:]
		}
		ans = append(ans, tmp)
	}
	return ans
}
