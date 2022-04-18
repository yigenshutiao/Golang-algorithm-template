package _146_lru_cache

//type LRUCache struct {
//	Capacity int
//	// hash表
//	Cache      map[int]*DLinkedNode
//	Head, Tail *DLinkedNode
//}
//
//// DLinkedNode 双向链表
//type DLinkedNode struct {
//	Key, Value int
//	Prev, Next *DLinkedNode
//}
//
//// 初始化一个双链表节点
//func initDLinkedNode(key, value int) *DLinkedNode {
//	return &DLinkedNode{Key: key, Value: value}
//}
//
//func Constructor(capacity int) LRUCache {
//	l := LRUCache{
//		Cache:    map[int]*DLinkedNode{},
//		Head:     initDLinkedNode(0, 0),
//		Tail:     initDLinkedNode(0, 0),
//		Capacity: capacity,
//	}
//	l.Head.Next = l.Tail
//	l.Tail.Prev = l.Head
//	return l
//}
//
//func (c *LRUCache) Get(key int) int {
//	// 去map里查找，不存在返回 -1
//	if _, ok := c.Cache[key]; !ok {
//		return -1
//	}
//	// 存在时，把这个节点提到头部
//	node := c.Cache[key]
//	c.moveToHead(node)
//	// 返回节点的值
//	return node.Value
//}
//
//// Put ***要记得更新cache里面的值和删除node里面的值***
//func (c *LRUCache) Put(key int, value int) {
//	// 看lru里有没有这个key
//	if _, ok := c.Cache[key]; !ok { // 如果没有
//		// 初始化这个节点
//		node := initDLinkedNode(key, value)
//		// 把节点添加到map里面
//		c.Cache[key] = node
//		// 把节点提到双向链表头部
//		c.addToHead(node)
//		// 如果size 大于 cap
//		if len(c.Cache) > c.Capacity {
//			// 把尾部节点删掉，这个函数应返回此节点
//			removed := c.removeTail()
//			// map删掉尾部节点
//			delete(c.Cache, removed.Key)
//		}
//	} else { // 如果有
//		// 获取这个key
//		node := c.Cache[key]
//		// 对于这个key，赋值新的value
//		node.Value = value
//		// 把这个节点移到双向链表头部
//		c.moveToHead(node)
//	}
//}
//
//// 这个是最base的函数，把新节点加到头节点的下一位
//func (c *LRUCache) addToHead(node *DLinkedNode) {
//	node.Prev = c.Head
//	node.Next = c.Head.Next
//	// 头结点下一个节点的prev指向node
//	c.Head.Next.Prev = node
//	// 头结点的next指向node
//	c.Head.Next = node
//}
//
//// 这个是最base的函数，删除当前节点
//func (c *LRUCache) removeNode(node *DLinkedNode) {
//	node.Prev.Next = node.Next
//	node.Next.Prev = node.Prev
//}
//
//// 把当前节点移动到头结点的下一个
//func (c *LRUCache) moveToHead(node *DLinkedNode) {
//	// 删掉链表当前的位置
//	c.removeNode(node)
//	// 把这个节点移到链表头部
//	c.addToHead(node)
//}
//
//// 把尾节点前一个节点删掉，并返回其值
//func (c *LRUCache) removeTail() *DLinkedNode {
//	node := c.Tail.Prev
//	c.removeNode(node)
//	return node
//}
