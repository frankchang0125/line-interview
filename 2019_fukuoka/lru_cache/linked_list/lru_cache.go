package lrucache

// Reference: http://bit.ly/2UE2VCX
type LRUCache struct {
	CacheMap  map[int]*CacheNode
	CacheHead *CacheNode
	CacheTail *CacheNode
	CacheSize int
}

// LRU node is appended/moved to the head of list.
type CacheNode struct {
	Key  int
	Val  int
	Prev *CacheNode
	Next *CacheNode
}

func Constructor() LRUCache {
	return LRUCache{
		CacheMap:  map[int]*CacheNode{},
		CacheHead: nil,
		CacheTail: nil,
		CacheSize: 0,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.CacheMap[key]
	if !ok {
		// Node not found.
		return -1
	}

	// Node found, move node to the head of list.
	if node != this.CacheHead {
		this.MoveNodeToHead(node)
	}

	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.CacheMap[key]
	if !ok {
		node := &CacheNode{Key: key, Val: value}
		this.AddNewNode(node)
		this.CacheMap[key] = node
		this.CacheSize++
	} else {
		// Node key found in cache,
		// move the node to the head of list.
		if node != this.CacheHead {
			this.MoveNodeToHead(node)
		}

		// Node value may be different, update node value.
		node.Val = value
	}
}

func (this *LRUCache) Evict() {
	if this.CacheSize > 0 {
		node := this.CacheTail

		if this.CacheSize == 1 {
			this.CacheHead = nil
			this.CacheTail = nil
		} else {
			this.CacheTail = this.CacheTail.Prev
		}

		delete(this.CacheMap, node.Key)
		this.CacheSize--
	}
}

func (this *LRUCache) Remove(key int) int {
	node, ok := this.CacheMap[key]
	if !ok {
		return -1
	}

	if this.CacheSize == 1 {
		this.CacheHead = nil
		this.CacheTail = nil
	} else if node == this.CacheHead {
		this.CacheHead = this.CacheHead.Next
	} else if node == this.CacheTail {
		this.CacheTail = this.CacheTail.Prev
	} else {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
	}

	delete(this.CacheMap, node.Key)
	this.CacheSize--
	return node.Val
}

func (this *LRUCache) RemoveOutdatedNode() {
	delete(this.CacheMap, this.CacheTail.Key)
	prev := this.CacheTail.Prev
	if prev != nil {
		prev.Next = nil
		this.CacheTail = prev
	}
}

func (this *LRUCache) AddNewNode(node *CacheNode) {
	if this.CacheSize == 0 {
		this.CacheHead = node
		this.CacheTail = node
	} else {
		this.CacheHead.Prev = node
		node.Next = this.CacheHead
		this.CacheHead = node
	}
}

func (this *LRUCache) MoveNodeToHead(node *CacheNode) {
	if this.CacheSize == 0 {
		this.CacheHead = node
		this.CacheTail = node
	} else {
		if node == this.CacheTail {
			this.CacheTail = node.Prev
		}

		if node.Prev != nil {
			node.Prev.Next = node.Next
		}

		if node.Next != nil {
			node.Next.Prev = node.Prev
		}

		node.Prev = nil
		node.Next = this.CacheHead
		this.CacheHead.Prev = node
		this.CacheHead = node
	}
}
