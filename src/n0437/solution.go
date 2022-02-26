package n0437

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) (ans int) {
	if root == nil {
		return
	}
	preSum := make(map[int]int)
	preSum[0] = 1
	var dfs func(node *TreeNode, sumOfPath int)
	dfs = func(node *TreeNode, sumOfPath int) {
		if node == nil {
			return
		}
		sumOfPath += node.Val
		prePathNum := preSum[sumOfPath-targetSum]
		ans += prePathNum
		preSum[sumOfPath]++
		dfs(node.Left, sumOfPath)
		dfs(node.Right, sumOfPath)
		preSum[sumOfPath]--
	}
	dfs(root, 0)
	return
}

// func rootSum(root *TreeNode, targetSum int) (res int) {
//     if root == nil {
//         return
//     }
//     val := root.Val
//     if val == targetSum {
//         res++
//     }
//     res += rootSum(root.Left, targetSum-val)
//     res += rootSum(root.Right, targetSum-val)
//     return
// }

// func pathSum(root *TreeNode, targetSum int) int {
//     if root == nil {
//         return 0
//     }
//     res := rootSum(root, targetSum)
//     res += pathSum(root.Left, targetSum)
//     res += pathSum(root.Right, targetSum)
//     return res
// }
