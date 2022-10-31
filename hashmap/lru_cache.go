package hashmap

// LRUCache
// @summary: LeetCode #146 (Medium); CodeTop MS
// @author : binqibang
// @created: 2022/8/26 10:52
type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*Node
	head, tail *Node
}

// Node for double linked list.
type Node struct {
	key, value int
	prev, next *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{key: 0, value: 0}
	tail := &Node{key: 0, value: 0}
	head.next = tail
	tail.prev = head
	return LRUCache{
		size:     0,
		capacity: capacity,
		head:     head,
		tail:     tail,
		cache:    map[int]*Node{},
	}
}

func (lru *LRUCache) Get(key int) int {
	if _, ok := lru.cache[key]; !ok {
		return -1
	}
	node := lru.cache[key]
	lru.moveFirst(node)
	return node.value
}

func (lru *LRUCache) Put(key int, value int) {
	if _, ok := lru.cache[key]; !ok {
		node := &Node{key: key, value: value}
		lru.cache[key] = node
		lru.addFirst(node)
		lru.size++
		if lru.size > lru.capacity {
			lru.removeLast()
			lru.size--
		}
	} else {
		node := lru.cache[key]
		node.value = value
		lru.moveFirst(node)
	}
}

func (lru *LRUCache) addFirst(node *Node) {
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node
	lru.head.next = node
}

func (lru *LRUCache) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (lru *LRUCache) moveFirst(node *Node) {
	lru.remove(node)
	lru.addFirst(node)
}

func (lru *LRUCache) removeLast() {
	node := lru.tail.prev
	lru.remove(node)
	delete(lru.cache, node.key)
}
