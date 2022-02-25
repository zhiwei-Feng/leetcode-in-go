package n0114

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	rightNode := root.Right
	leftNode := root.Left
	root.Right, root.Left = nil, nil
	flatten(leftNode)
	flatten(rightNode)
	if leftNode == nil {
		root.Right = rightNode
	} else {
		root.Right = leftNode
		p := root.Right
		for p.Right != nil {
			p = p.Right
		}
		p.Right = rightNode
	}
}
