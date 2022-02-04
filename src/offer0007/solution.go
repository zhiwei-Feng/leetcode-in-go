package offer0007

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{preorder[0], nil, nil}
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			leftTreeNum := len(inorder[:i])
			root.Left = buildTree(preorder[1:leftTreeNum+1], inorder[:i])
			root.Right = buildTree(preorder[leftTreeNum+1:], inorder[i+1:])
			break
		}
	}
	return root
}
