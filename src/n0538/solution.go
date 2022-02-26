package n0538

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	dfs(root, 0)
	return root
}

func dfs(node *TreeNode, prevSum int) int {
	if node == nil {
		return prevSum
	}

	rightSum := dfs(node.Right, prevSum)
	node.Val += rightSum
	return dfs(node.Left, node.Val)
}
