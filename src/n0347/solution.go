package n0347

type Node struct {
	Freq       int
	Val        int
	Prev, Next *Node
}

type List struct {
	Head, Tail *Node
}

func NewList() List {
	newList := List{
		Head: &Node{},
		Tail: &Node{},
	}
	newList.Head.Next = newList.Tail
	newList.Tail.Prev = newList.Head
	return newList
}
func IsEmpty(list List) bool {
	return list.Head == list.Tail
}
func RemoveNode(node *Node) {
	prev := node.Prev
	next := node.Next
	prev.Next = next
	next.Prev = prev
}
func (l List) AddToHead(node *Node) {
	next := l.Head.Next
	next.Prev = node
	node.Next = next
	l.Head.Next = node
	node.Prev = l.Head
}
func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]List)
	maxFreq := 0
	nodeMap := make(map[int]*Node)
	for _, v := range nums {
		if vNode, ok := nodeMap[v]; !ok {
			newNode := &Node{Val: v, Freq: 1}
			nodeMap[v] = newNode
			if _, listOk := freqMap[1]; !listOk {
				freqMap[1] = NewList()
			}
			freqMap[1].AddToHead(newNode)
			maxFreq=1
		} else {
			RemoveNode(vNode)
			vNode.Freq++
			if fList, listOk := freqMap[vNode.Freq]; !listOk {
				freqMap[vNode.Freq] = NewList()
			} else {
				fList.AddToHead(vNode)
			}
			if vNode.Freq>maxFreq{
				maxFreq=vNode.Freq
			}
		}
	}

	count := 0
	freq := maxFreq
	ans := []int{}
	for count < k {
		if v, ok := freqMap[freq]; !ok || IsEmpty(v) {
			freq--
			continue
		} else {
			p := v.Head
			for p != v.Tail && count < k {
				ans = append(ans, p.Val)
				count++
				p = p.Next
			}
			if count == k {
				break
			}
			freq--
		}
	}
	return ans
}
