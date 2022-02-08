package offer0037

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	// write code here
	var ans strings.Builder
	// type pair struct {
	// 	node *TreeNode
	// 	num  int
	// }
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		curNode := queue[0]
		queue = queue[1:]
		if curNode == nil {
			ans.WriteString("#")
			ans.WriteString(" ")
			continue
		}
		ans.WriteString(strconv.Itoa(curNode.Val))
		ans.WriteString(" ")

		queue = append(queue, curNode.Left)
		queue = append(queue, curNode.Right)
	}

	ret := ans.String()
	return ret[:len(ret)-1]
}
func Deserialize(s string) *TreeNode {
	// write code here
	if s == "" {
		return nil
	}
	queue := []*TreeNode{}

	vals := strings.Split(s, " ")
	rootVal, _ := strconv.Atoi(vals[0])
	root := &TreeNode{rootVal, nil, nil}
	queue = append(queue, root)
	i := 1
	for len(queue) != 0 {
		curNode := queue[0]
		queue = queue[1:]
		if vals[i] != "#" {
			x, _ := strconv.Atoi(vals[i])
			curNode.Left = &TreeNode{x, nil, nil}
			queue = append(queue, curNode.Left)
		}
		i++
		if vals[i] != "#" {
			x, _ := strconv.Atoi(vals[i])
			curNode.Right = &TreeNode{x, nil, nil}
			queue = append(queue, curNode.Right)
		}
		i++
	}

	return root
}
