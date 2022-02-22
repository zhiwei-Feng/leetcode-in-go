package n0106

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	var build func(inS, inE, postS, postE int) *TreeNode
	build = func(inS, inE, postS, postE int) *TreeNode {
		if postS>postE {
			return nil
		}
		if postS==postE {
			return &TreeNode{Val: postorder[postE]}
		}
		root := &TreeNode{Val: postorder[postE]}
		rootInOrderInd := inS
		for ;rootInOrderInd<=inE;rootInOrderInd++{
			if inorder[rootInOrderInd]==root.Val {
				break
			}
		}
		leftNum := rootInOrderInd-inS
		root.Left = build(inS, rootInOrderInd-1, postS, postS+leftNum-1)
		root.Right = build(rootInOrderInd+1, inE, postS+leftNum, postE-1)
		
		return root
	}

	return build(0, len(inorder)-1, 0, len(postorder)-1)
}