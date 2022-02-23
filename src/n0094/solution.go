package n0094

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var (
		ans     = make([]int, 0)
		inOrder func(node *TreeNode)
	)

	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		ans = append(ans, node.Val)
		inOrder(node.Right)
	}

	inOrder(root)
	return ans
}
