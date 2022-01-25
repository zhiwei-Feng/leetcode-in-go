package offer0036

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Convert(pRootOfTree *TreeNode) *TreeNode {
	if pRootOfTree == nil {
		return nil
	}

	var (
		head = &TreeNode{}
		pre  *TreeNode
		dfs  func(node *TreeNode)
	)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != nil {
			pre.Right = node
		}else {
			head = node
		}
		node.Left = pre
		pre = node
		dfs(node.Right)
	}

	dfs(pRootOfTree)
	head.Left = pre
	pre.Right = head
	return head
}
