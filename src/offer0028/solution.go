package offer0028

func isSymmetric(root *TreeNode) bool {
	return root == nil || recur(root.Left, root.Right)
}

func recur(leftNode, rightNode *TreeNode) bool {
	if leftNode == nil && rightNode == nil {
		return true
	}
	if leftNode != nil && rightNode != nil {
		return leftNode.Val == rightNode.Val &&
			recur(leftNode.Left, rightNode.Right) &&
			recur(leftNode.Right, rightNode.Left)
	}
	return false
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}
