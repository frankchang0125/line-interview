package lrucache

import (
	"container/list"
)

// Reference: http://bit.ly/2UE2VCX
type LRUCache struct {
	CacheMap  map[int]*list.Element
	CacheList *list.List
}

// LRU node is appended/moved to the head of list.
type CacheNode struct {
	Key  int
	Val  int
}

func Constructor() LRUCache {
	return LRUCache{
		CacheMap:  map[int]*list.Element{},
		CacheList: list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	element, ok := this.CacheMap[key]
	if !ok {
		// Node not found.
		return -1
	}

	// Node found, move node to the head of list.
	this.MoveNodeToHead(element)
	node := element.Value.(*CacheNode)
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	element, ok := this.CacheMap[key]
	if !ok {
		node := &CacheNode{Key: key, Val: value}
		element := this.AddNewNode(node)
		this.CacheMap[key] = element
	} else {
		// Node key found in cache,
		// move the node to the head of list.
		this.MoveNodeToHead(element)

		// Node value may be different, update node value.
		element.Value.(*CacheNode).Val = value
	}
}

func (this *LRUCache) Evict() {
	if this.CacheList.Len() > 0 {
		element := this.CacheList.Back()
		delete(this.CacheMap, element.Value.(*CacheNode).Key)
		this.CacheList.Remove(element)
	}
}

func (this *LRUCache) Remove(key int) int {
	element, ok := this.CacheMap[key]
	if !ok {
		return -1
	}
	delete(this.CacheMap, key)
	this.CacheList.Remove(element)
	return element.Value.(*CacheNode).Val
}

func (this *LRUCache) RemoveOutdatedNode() {
	element := this.CacheList.Back()
	this.CacheList.Remove(element)
	node := element.Value.(*CacheNode)
	delete(this.CacheMap, node.Key)
}

func (this *LRUCache) AddNewNode(node *CacheNode) (*list.Element) {
	return this.CacheList.PushFront(node)
}

func (this *LRUCache) MoveNodeToHead(element *list.Element) {
	this.CacheList.MoveToFront(element)
}
