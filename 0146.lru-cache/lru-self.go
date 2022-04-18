package _146_lru_cache

type LRUCache struct {
	head *Node
	tail *Node
	info map[int]*Node
	capa int
}

type Node struct {
	key  int
	val  int
	pre  *Node
	next *Node
}

func Constructor(capacity int) LRUCache {
	info := make(map[int]*Node, capacity)
	lru := LRUCache{
		capa: capacity,
		head: &Node{},
		tail: &Node{},
		info: info,
	}

	lru.head.next = lru.tail
	lru.tail.pre = lru.head
	return lru
}

func (cache *LRUCache) Get(key int) int {
	if _, ok := cache.info[key]; !ok {
		// 如果不存在节点，返回-1
		return -1
	}
	// 如果存在节点, 返回值，把当前位置删除，并把节点放在头部
	node := cache.info[key]

	cache.removeNode(node)

	cache.addToHead(node)

	return node.val
}

func (cache *LRUCache) addToHead(node *Node) {
	node.pre = cache.head
	node.next = cache.head.next
	cache.head.next.pre = node
	cache.head.next = node
}

func (cache *LRUCache) removeNode(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (cache *LRUCache) Put(key int, value int) {
	if _, ok := cache.info[key]; ok {
		// 如果存在节点，则更新, 并放到头部
		node := cache.info[key]
		node.val = value

		// 把当前节点删掉
		cache.removeNode(node)

		// head的下一个节点置为这个节点
		cache.addToHead(node)

	} else {
		// 如果不存在节点，则插入到头部
		node := &Node{key: key, val: value}
		cache.info[key] = node
		// 把这个节点插入到头部
		cache.addToHead(node)

		// 看map的长度是否大于cap, 如果大于需要把尾部节点干掉
		if len(cache.info) > cache.capa {
			nodePreTail := cache.tail.pre

			cache.removeNode(nodePreTail)

			delete(cache.info, nodePreTail.key)
		}
	}
}
