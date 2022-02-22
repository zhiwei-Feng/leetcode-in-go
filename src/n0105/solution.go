package n0105

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	var build func(preS, preE int, inS, inE int) *TreeNode
	build = func(preS, preE, inS, inE int) *TreeNode {
		if preS > preE {
			return nil
		}
		if preS == preE {
			return &TreeNode{Val: preorder[preS]}
		}
		root := &TreeNode{Val: preorder[preS]}
		rootInOrderInd := inS
		for ; rootInOrderInd <= inE; rootInOrderInd++ {
			if inorder[rootInOrderInd] == root.Val {
				break
			}
		}
		leftNum := rootInOrderInd - inS
		root.Left = build(preS+1, preS+leftNum, inS, rootInOrderInd-1)
		root.Right = build(preS+leftNum+1, preE, rootInOrderInd+1, inE)
		return root
	}

	return build(0, len(preorder)-1, 0, len(inorder)-1)
}
