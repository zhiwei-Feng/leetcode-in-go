package n0124

import "math"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	var maxGain func(node *TreeNode) int
	var ans = math.MinInt

	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		priceNow := node.Val + leftGain + rightGain
		if priceNow > ans {
			ans = priceNow
		}

		return node.Val + max(leftGain, rightGain)
	}

	maxGain(root)
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
