package offer0054

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	var (
		count = 0
		dfs   func(node *TreeNode)
		ans   int
	)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		count++
		if count == k {
			ans = node.Val
			return
		}
		dfs(node.Left)
	}

	dfs(root)
	return ans
}
