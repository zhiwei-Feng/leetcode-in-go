package n0543

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var getH func(node *TreeNode) int
	ans := 0
	getH = func(node *TreeNode) int {
		if node.Left == nil && node.Right == nil {
			return 0
		}

		leftH := 0
		rightH := 0
		curPath := 0
		if node.Left != nil {
			leftH = getH(node.Left)
			curPath += leftH + 1
		}
		if node.Right != nil {
			rightH = getH(node.Right)
			curPath += rightH + 1
		}
		if ans < curPath {
			ans = curPath
		}
		return max(leftH, rightH) + 1
	}

	getH(root)
	return ans
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
