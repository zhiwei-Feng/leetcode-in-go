package n0863

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	father := make(map[int]*TreeNode)
	ans := []int{}
	var dfs func(*TreeNode, *TreeNode)
	dfs = func(node *TreeNode, fa *TreeNode) {
		father[node.Val] = fa
		if node.Left != nil {
			dfs(node.Left, node)
		}
		if node.Right != nil {
			dfs(node.Right, node)
		}
	}
	dfs(root, nil)

	var searchDis func(*TreeNode, int, int)
	searchDis = func(current *TreeNode, from int, distance int) {
		if distance == k {
			ans = append(ans, current.Val)
			return
		}
		curVal := current.Val
		if father[curVal] != nil && father[curVal].Val != from {
			searchDis(father[curVal], curVal, distance+1)
		}
		if current.Left != nil && current.Left.Val != from {
			searchDis(current.Left, curVal, distance+1)
		}
		if current.Right != nil && current.Right.Val != from {
			searchDis(current.Right, curVal, distance+1)
		}
	}
	searchDis(target, -1, 0)
	return ans
}
