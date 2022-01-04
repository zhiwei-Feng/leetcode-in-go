package n0146

type LinkNode struct {
	Val        int
	Key        int
	Prev, Next *LinkNode
}

type LRUCache struct {
	capacity  int
	QueueHead *LinkNode
	QueueTail *LinkNode
	NodeMap   map[int]*LinkNode
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		capacity:  capacity,
		QueueHead: &LinkNode{},
		QueueTail: &LinkNode{},
		NodeMap:   make(map[int]*LinkNode),
	}
	cache.QueueHead.Next = cache.QueueTail
	cache.QueueTail.Prev = cache.QueueHead
	return cache
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.NodeMap[key]; !ok {
		return -1
	} else {
		this.MoveToHead(v)
		return v.Val
	}
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.NodeMap[key]; !ok {
		node := &LinkNode{Key: key, Val: value}
		this.NodeMap[key] = node
		this.AddToHead(node)
		if len(this.NodeMap) > this.capacity {
			delTail := this.RemoveTail()
			delete(this.NodeMap, delTail)
		}
	} else {
		v.Val = value
		this.MoveToHead(v)
	}
}

func (this *LRUCache) MoveToHead(node *LinkNode) {
	this.RemoveNode(node)
	this.AddToHead(node)
}

func (this *LRUCache) RemoveNode(node *LinkNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) AddToHead(node *LinkNode) {
	node.Prev = this.QueueHead
	node.Next = this.QueueHead.Next
	this.QueueHead.Next.Prev = node
	this.QueueHead.Next = node
}

func (this *LRUCache) RemoveTail() int {
	delTail := this.QueueTail.Prev
	newTail := this.QueueTail.Prev.Prev
	newTail.Next = this.QueueTail
	this.QueueTail.Prev = newTail
	return delTail.Key
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
