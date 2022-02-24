package n0617

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil || root2 == nil {
		if root1 == nil {
			return root2
		}
		return root1
	}
	newRoot := &TreeNode{Val: root1.Val + root2.Val}
	newRoot.Left = mergeTrees(root1.Left, root2.Left)
	newRoot.Right = mergeTrees(root1.Right, root2.Right)
	return newRoot
}