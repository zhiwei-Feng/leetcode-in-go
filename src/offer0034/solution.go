package offer0034

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := [][]int{}
	var dfs func(node *TreeNode, target int, tmpPath []int)
	dfs = func(node *TreeNode, target int, tmpPath []int) {
		if node.Left == nil && node.Right == nil {
			// leaf node
			if node.Val == target {
				p := make([]int, len(tmpPath))
				copy(p, tmpPath)
				p = append(p, node.Val)
				ans = append(ans, p)
			}
			return
		}
		if node.Left != nil {
			tmpPath = append(tmpPath, node.Val)
			dfs(node.Left, target-node.Val, tmpPath)
			tmpPath = tmpPath[:len(tmpPath)-1]
		}
		if node.Right != nil {
			tmpPath = append(tmpPath, node.Val)
			dfs(node.Right, target-node.Val, tmpPath)
			tmpPath = tmpPath[:len(tmpPath)-1]
		}
	}

	dfs(root, target, []int{})
	return ans
}
