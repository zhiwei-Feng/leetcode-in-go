package n0460

type LinkNode struct {
	Key        int
	Val        int
	Freq       int
	Prev, Next *LinkNode
}

type LinkList struct {
	Head, Tail *LinkNode
}

func (list *LinkList) AddToHead(node *LinkNode) {
	node.Next = list.Head.Next
	node.Prev = list.Head
	list.Head.Next.Prev = node
	list.Head.Next = node
}

func (list *LinkList) RemoveTheTail() *LinkNode {
	if list.Head.Next == list.Tail {
		return nil
	}
	delNode := list.Tail.Prev
	list.Tail.Prev.Prev.Next = list.Tail
	list.Tail.Prev = list.Tail.Prev.Prev
	return delNode
}

func (list *LinkList) IsEmpty() bool {
	return list.Head.Next == list.Tail
}

type LFUCache struct {
	Capacity int
	MinFreq  int
	FreqMap  map[int]*LinkList
	KeyMap   map[int]*LinkNode
}

func Constructor(capacity int) LFUCache {
	cache := LFUCache{
		Capacity: capacity,
		MinFreq:  0,
		FreqMap:  make(map[int]*LinkList),
		KeyMap:   make(map[int]*LinkNode),
	}
	return cache
}

func NewList() *LinkList {
	newList := &LinkList{
		Head: &LinkNode{},
		Tail: &LinkNode{},
	}
	newList.Head.Next = newList.Tail
	newList.Tail.Prev = newList.Head
	return newList
}

func (this *LFUCache) Get(key int) int {
	if this.Capacity==0 {
		return -1
	}
	node, ok := this.KeyMap[key]
	if !ok {
		return -1
	}
	val, freq := node.Val, node.Freq
	RemoveNode(node)
	if this.FreqMap[freq].IsEmpty() {
		delete(this.FreqMap, freq)
		if this.MinFreq == freq {
			this.MinFreq += 1
		}
	}
	list, exist := this.FreqMap[freq+1]
	if !exist {
		list = NewList()
	}
	newNode := &LinkNode{key, val, freq + 1, nil, nil}
	list.AddToHead(newNode)
	this.FreqMap[freq+1] = list
	this.KeyMap[key] = newNode
	return val
}

func (this *LFUCache) Put(key int, value int) {
	if this.Capacity==0 {
		return
	}
	if node, ok := this.KeyMap[key]; !ok {
		if len(this.KeyMap) == this.Capacity {
			delNode := this.FreqMap[this.MinFreq].RemoveTheTail()
			delete(this.KeyMap, delNode.Key)
			if this.FreqMap[this.MinFreq].IsEmpty() {
				delete(this.FreqMap, this.MinFreq)
			}
		}

		list, exist := this.FreqMap[1]
		if !exist {
			list = NewList()
		}
		newNode := &LinkNode{key, value, 1, nil, nil}
		list.AddToHead(newNode)
		this.FreqMap[1] = list
		this.KeyMap[key] = newNode
		this.MinFreq = 1
	} else {
		freq := node.Freq
		RemoveNode(node)
		if this.FreqMap[freq].IsEmpty() {
			delete(this.FreqMap, freq)
			if this.MinFreq == freq {
				this.MinFreq += 1
			}
		}
		list, exist := this.FreqMap[freq+1]
		if !exist {
			list = NewList()
		}
		newNode := &LinkNode{key, value, freq + 1, nil, nil}
		list.AddToHead(newNode)
		this.FreqMap[freq+1] = list
		this.KeyMap[key] = newNode
	}
}

func RemoveNode(node *LinkNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
