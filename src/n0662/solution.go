package n0662

/**
 * Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	q := []State{{root, 1}}
	ans := 1
	for len(q) > 0 {
		size := len(q)
		minId := math.MaxInt64
		maxId := math.MinInt64
		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]
			if cur.Node.Left != nil {
				q = append(q, State{cur.Node.Left, cur.Id * 2})
			}
			if cur.Node.Right != nil {
				q = append(q, State{cur.Node.Right, cur.Id*2 + 1})
			}
			minId = min(minId, cur.Id)
			maxId = max(maxId, cur.Id)
		}
		ans = max(ans, maxId-minId+1)
	}
	return ans
}

type State struct {
	Node *TreeNode
	Id   int
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
