package n0102

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	var lastLevelNode *TreeNode = root
	var candLastLevelNode *TreeNode
	var tmpList = []int{}
	q := []*TreeNode{root}
	for len(q) != 0 {
		curNode := q[0]
		q = q[1:]
		tmpList = append(tmpList, curNode.Val)
		if curNode.Left != nil {
			q = append(q, curNode.Left)
			candLastLevelNode = curNode.Left
		}
		if curNode.Right != nil {
			q = append(q, curNode.Right)
			candLastLevelNode = curNode.Right
		}
		if curNode == lastLevelNode {
			lastLevelNode = candLastLevelNode
			t := make([]int, len(tmpList))
			copy(t, tmpList)
			ans = append(ans, t)
			tmpList = []int{}
		}
	}
	return ans
}
