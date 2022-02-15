package _146_lru_cache

type cache struct {
	size int

	data map[string]*Node
	head *Node
	tail *Node
}

type Node struct {
	key  string
	val  string
	pre  *Node
	next *Node
}

func (c *cache) Get(key string) string {
	node, exist := c.moveToHead(key)
	if !exist {
		return ""
	}
	return node.val
}

func (c *cache) Set(key, val string) {
	node, exist := c.moveToHead(key)
	if exist {
		node.val = val
		return
	}

	if len(c.data) >= c.size {
		c.disuse()
	}

	node = &Node{
		key: key,
		val: val,
	}
	c.data[key] = node

	node.next = c.head
	if c.head != nil {
		c.head.pre = node
	} else {
		c.tail = node
	}
	c.head = node
}

func (c *cache) disuse() {
	delete(c.data, c.tail.key)
	c.tail.pre.next = nil
	c.tail = c.tail.pre
}

func (c *cache) moveToHead(key string) (*Node, bool) {
	node, ok := c.data[key]
	if !ok {
		return nil, false
	}

	if c.tail == node {
		c.tail = node.pre
	}

	if node.pre != nil {
		node.pre.next = node.next
	}
	if node.next != nil {
		node.next.pre = node.pre
	}

	node.pre = nil
	node.next = c.head
	c.head.pre = node
	c.head = node
	return node, true
}
