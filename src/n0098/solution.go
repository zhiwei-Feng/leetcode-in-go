package n0098

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var inOrderPrevNode *TreeNode
	var inOrder func(node *TreeNode) bool
	inOrder = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if !inOrder(node.Left) {
			return false
		}
		if inOrderPrevNode!=nil && node.Val <= inOrderPrevNode.Val {
			return false
		}
		inOrderPrevNode = node
		return inOrder(node.Right)
	}

	return inOrder(root)
}
