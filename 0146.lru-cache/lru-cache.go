package _146_lru_cache

type LRUCache struct {
	size     int
	capacity int
	// hash表
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

// DLinkedNode 双向链表
type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

// 初始化一个双链表节点
func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*DLinkedNode{},
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (c *LRUCache) Get(key int) int {
	// 去map里查找，不存在返回 -1
	if _, ok := c.cache[key]; !ok {
		return -1
	}
	// 存在时，把这个节点提到头部
	node := c.cache[key]
	c.moveToHead(node)
	// 返回节点的值
	return node.value
}

func (c *LRUCache) Put(key int, value int) {
	// 看lru里有没有这个key
	if _, ok := c.cache[key]; !ok { // 如果没有
		// 初始化这个节点
		node := initDLinkedNode(key, value)
		// 把节点添加到map里面
		c.cache[key] = node
		// 把节点提到双向链表头部
		c.addToHead(node)
		// 链表长度++
		c.size++
		// 如果size 大于 cap
		if c.size > c.capacity {
			// 把尾部节点删掉，这个函数应返回此节点
			removed := c.removeTail()
			// map删掉尾部节点
			delete(c.cache, removed.key)
			// 双向链表size减1
			c.size--
		}
	} else { // 如果有
		// 获取这个key
		node := c.cache[key]
		// 对于这个key，赋值新的value
		node.value = value
		// 把这个节点移到双向链表头部
		c.moveToHead(node)
	}
}

// 这个是最base的函数，把新节点加到头节点的下一位
func (c *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = c.head
	node.next = c.head.next
	// 头结点下一个节点的prev指向node
	c.head.next.prev = node
	// 头结点的next指向node
	c.head.next = node
}

// 这个是最base的函数，删除当前节点
func (c *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 把当前节点移动到头结点的下一个
func (c *LRUCache) moveToHead(node *DLinkedNode) {
	// 删掉链表当前的位置
	c.removeNode(node)
	// 把这个节点移到链表头部
	c.addToHead(node)
}

// 把尾节点前一个节点删掉，并返回其值
func (c *LRUCache) removeTail() *DLinkedNode {
	node := c.tail.prev
	c.removeNode(node)
	return node
}
