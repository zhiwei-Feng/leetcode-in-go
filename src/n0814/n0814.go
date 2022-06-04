package n0814

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		var isLeftIncludeOne = dfs(node.Left)
		var isRightIncludeOne = dfs(node.Right)
		if !isLeftIncludeOne {
			// removed
			node.Left = nil
		}
		if !isRightIncludeOne {
			// removed
			node.Right = nil
		}
		return !(node.Left == nil && node.Right == nil && node.Val == 0)
	}

	var res = dfs(root)
	if !res {
		return nil
	}
	return root
}
