package _146_lru_cache

//type cache struct {
//	size int
//
//	data map[string]*Node
//	Head *Node
//	Tail *Node
//}
//
//type Node struct {
//	Key  string
//	Val  string
//	Pre  *Node
//	Next *Node
//}
//
//func (c *cache) Get(Key string) string {
//	node, exist := c.moveToHead(Key)
//	if !exist {
//		return ""
//	}
//	return node.Val
//}
//
//func (c *cache) Set(Key, Val string) {
//	node, exist := c.moveToHead(Key)
//	if exist {
//		node.Val = Val
//		return
//	}
//
//	if len(c.data) >= c.size {
//		c.disuse()
//	}
//
//	node = &Node{
//		Key: Key,
//		Val: Val,
//	}
//	c.data[Key] = node
//
//	node.Next = c.Head
//	if c.Head != nil {
//		c.Head.Pre = node
//	} else {
//		c.Tail = node
//	}
//	c.Head = node
//}
//
//func (c *cache) disuse() {
//	delete(c.data, c.Tail.Key)
//	c.Tail.Pre.Next = nil
//	c.Tail = c.Tail.Pre
//}
//
//func (c *cache) moveToHead(Key string) (*Node, bool) {
//	node, ok := c.data[Key]
//	if !ok {
//		return nil, false
//	}
//
//	if c.Tail == node {
//		c.Tail = node.Pre
//	}
//
//	if node.Pre != nil {
//		node.Pre.Next = node.Next
//	}
//	if node.Next != nil {
//		node.Next.Pre = node.Pre
//	}
//
//	node.Pre = nil
//	node.Next = c.Head
//	c.Head.Pre = node
//	c.Head = node
//	return node, true
//}
