package _146_lru_cache

type LRUCache struct {
	Head *Node
	Tail *Node
	Info map[int]*Node
	Capa int
}

type Node struct {
	Key  int
	Val  int
	Pre  *Node
	Next *Node
}

func Constructor(capacity int) LRUCache {
	info := make(map[int]*Node, capacity)
	lru := LRUCache{
		Capa: capacity,
		Head: &Node{},
		Tail: &Node{},
		Info: info,
	}

	lru.Head.Next = lru.Tail
	lru.Tail.Pre = lru.Head
	return lru
}

func (cache *LRUCache) Get(key int) int {
	if _, ok := cache.Info[key]; !ok {
		// 如果不存在节点，返回-1
		return -1
	}
	// 如果存在节点, 返回值，把当前位置删除，并把节点放在头部
	node := cache.Info[key]

	cache.removeNode(node)

	cache.addToHead(node)

	return node.Val
}

func (cache *LRUCache) addToHead(node *Node) {
	node.Pre = cache.Head
	node.Next = cache.Head.Next
	cache.Head.Next.Pre = node
	cache.Head.Next = node
}

func (cache *LRUCache) removeNode(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (cache *LRUCache) Put(key int, value int) {
	if _, ok := cache.Info[key]; ok {
		// 如果存在节点，则更新, 并放到头部
		node := cache.Info[key]
		node.Val = value

		// 把当前节点删掉
		cache.removeNode(node)

		// head的下一个节点置为这个节点
		cache.addToHead(node)

	} else {
		// 如果不存在节点，则插入到头部
		node := &Node{Key: key, Val: value}
		cache.Info[key] = node
		// 把这个节点插入到头部
		cache.addToHead(node)

		// 看map的长度是否大于cap, 如果大于需要把尾部节点干掉
		if len(cache.Info) > cache.Capa {
			nodePreTail := cache.Tail.Pre

			cache.removeNode(nodePreTail)

			delete(cache.Info, nodePreTail.Key)
		}
	}
}
