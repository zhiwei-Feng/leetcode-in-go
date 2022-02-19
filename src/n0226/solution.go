package n0226

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	revLeft := invertTree(root.Left)
	revRight := invertTree(root.Right)
	root.Left = revRight
	root.Right = revLeft
	return root
}
