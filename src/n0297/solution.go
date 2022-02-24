package n0297

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var builder strings.Builder
	builder.WriteByte('[')

	// 层序遍历插入法
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]
		if curNode == nil {
			builder.WriteString("null" + ",")
			continue
		}
		builder.WriteString(fmt.Sprintf("%d,", curNode.Val))
		queue = append(queue, curNode.Left)
		queue = append(queue, curNode.Right)
	}

	builder.WriteByte(']')
	return builder.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) <= 2 {
		return nil
	}
	content := data[1 : len(data)-2]
	treeVals := strings.Split(content, ",")
	if treeVals[0] == "null" {
		return nil
	}
	rootVal, _ := strconv.Atoi(treeVals[0])
	root := &TreeNode{Val: rootVal}
	queue := []*TreeNode{root}
	curPoint := 1
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]
		leftVal := treeVals[curPoint]
		RightVal := treeVals[curPoint+1]
		if leftVal != "null" {
			left, _ := strconv.Atoi(leftVal)
			leftNode := &TreeNode{Val: left}
			curNode.Left = leftNode
			queue = append(queue, leftNode)
		}
		if RightVal != "null" {
			right, _ := strconv.Atoi(RightVal)
			rightNode := &TreeNode{Val: right}
			curNode.Right = rightNode
			queue = append(queue, rightNode)
		}
		curPoint+=2
	}
	return root
}