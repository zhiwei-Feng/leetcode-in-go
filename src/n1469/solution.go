package n1469

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getLonelyNodes(root *TreeNode) []int {
	ans := []int{}
	var dfs func(p *TreeNode, isUniq bool)
	dfs = func(p *TreeNode, isUniq bool) {
		if isUniq {
			ans = append(ans, p.Val)
		}
		count := 0
		if p.Left != nil {
			count++
		}
		if p.Right != nil {
			count++
		}
		if p.Left != nil {
			if count == 2 {
				dfs(p.Left, false)
			} else {
				dfs(p.Left, true)
			}
		}
		if p.Right != nil {
			if count == 2 {
				dfs(p.Right, false)
			} else {
				dfs(p.Right, true)
			}
		}
	}

	dfs(root, false)
	return ans
}
