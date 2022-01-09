package offer0035

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	cached := make(map[*Node]*Node)
	return deepCopy(head, cached)
}

func deepCopy(node *Node, cachedNode map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	if v, ok := cachedNode[node]; ok {
		return v
	}

	newNode := &Node{
		Val: node.Val,
	}
	cachedNode[node] = newNode // 不能够放在next和random copy完的后面，会内存溢出
	newNode.Next = deepCopy(node.Next, cachedNode)
	newNode.Random = deepCopy(node.Random, cachedNode)
	return newNode
}
