package n0146

type LinkNode struct {
	Key, Val   int
	Prev, Next *LinkNode
}
type LRUCache struct {
	capacity           int
	KeyMap             map[int]*LinkNode
	LinkHead, LinkTail *LinkNode
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		capacity: capacity,
		KeyMap:   make(map[int]*LinkNode),
		LinkHead: &LinkNode{},
		LinkTail: &LinkNode{},
	}
	cache.LinkHead.Next = cache.LinkTail
	cache.LinkTail.Prev = cache.LinkHead
	return cache
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.KeyMap[key]; !ok {
		return -1
	} else {
		this.MoveToHead(v)
		return v.Val
	}
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.KeyMap[key]; ok {
		this.MoveToHead(v)
		v.Val = value
	} else {
		newNode := &LinkNode{Val: value, Key: key}
		this.AddToHead(newNode)
		this.KeyMap[key] = newNode
		if len(this.KeyMap) > this.capacity {
			delNode := this.RemoveTail()
			delete(this.KeyMap, delNode.Key)
		}
	}
}

func (this *LRUCache) RemoveNode(node *LinkNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) AddToHead(node *LinkNode) {
	node.Next = this.LinkHead.Next
	node.Prev = this.LinkHead
	this.LinkHead.Next = node
	node.Next.Prev = node
}

func (this *LRUCache) MoveToHead(node *LinkNode) {
	this.RemoveNode(node)
	this.AddToHead(node)
}

func (this *LRUCache) RemoveTail() *LinkNode {
	delNode := this.LinkTail.Prev
	newTail := delNode.Prev
	newTail.Next = this.LinkTail
	this.LinkTail.Prev = newTail
	return delNode
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
