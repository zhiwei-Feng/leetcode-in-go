package offer0027

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	newLeft := mirrorTree(root.Left)
	newRight := mirrorTree(root.Right)
	root.Left = newRight
	root.Right = newLeft
	return root
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}
