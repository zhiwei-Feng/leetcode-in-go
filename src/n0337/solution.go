package n0337

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type SubTreeStatus struct {
	Selected int
	NotSelected int
}

// f(o)表示o被选中的情况下子树的最大权值
// g(o)表示o不被选中的情况下子树的最大权值
func rob(root *TreeNode) int {
	ans := dfs(root)
	return max(ans.Selected, ans.NotSelected)
}

func dfs(root *TreeNode) SubTreeStatus {
	if root == nil {
		return SubTreeStatus{0,0}
	}

	sub := SubTreeStatus{}
	l := dfs(root.Left)
	r := dfs(root.Right)
	sub.Selected = l.NotSelected+r.NotSelected+root.Val
	sub.NotSelected = max(l.NotSelected, l.Selected) + max(r.NotSelected, r.Selected)
	return sub
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// // f(o)表示o被选中的情况下子树的最大权值
// // g(o)表示o不被选中的情况下子树的最大权值
// func rob(root *TreeNode) int {
// 	f := make(map[*TreeNode]int)
// 	g := make(map[*TreeNode]int)
// 	dfs(root, f, g)
// 	return max(f[root], g[root])
// }

// func dfs(root *TreeNode, f, g map[*TreeNode]int) {
// 	if root == nil {
// 		return
// 	}

// 	dfs(root.Left, f, g)
// 	dfs(root.Right, f, g)
// 	f[root] = g[root.Left] + g[root.Right] + root.Val
// 	g[root] = max(f[root.Left], g[root.Left]) + max(f[root.Right], g[root.Right])
// }

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }
