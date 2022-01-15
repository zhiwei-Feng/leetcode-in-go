package offer0026

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	q := []*TreeNode{}
	q = append(q, A)
	for len(q) != 0 {
		top := q[0]
		q = q[1:]
		if isSame(top, B) {
			return true
		}
		if top.Left != nil {
			q = append(q, top.Left)
		}
		if top.Right != nil {
			q = append(q, top.Right)
		}
	}

	return false
}

func isSame(A, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	if A.Val != B.Val {
		return false
	}
	return isSame(A.Left, B.Left) && isSame(A.Right, B.Right)
}
