package base

import "errors"

type Node struct {
	Key   string
	Value string
	Prev  *Node
	Next  *Node
}

type LruCache struct {
	size     int
	capacity int
	Head     *Node
	Tail     *Node
}

func NewLruCache(capacity int) (*LruCache, error) {
	if capacity < 1 {
		return nil, errors.New("capacity must be greater than zero")
	}
	return &LruCache{
		capacity: capacity,
	}, nil
}

func (c *LruCache) Put(key string, value string) {
	node := c.find(key)
	if node != nil {
		node.Value = value
		return
	}
	c.trimPut(key, value)
}

func (c *LruCache) Get(key string) *string {
	node := c.find(key)
	if node == nil {
		return nil
	}
	res := &node.Value
	c.reorder(node)
	return res
}

func (c *LruCache) find(key string) *Node {
	if c.size == 0 {
		return nil
	}
	res := c.Head
	for i := 0; i < c.size; i++ {
		if res.Key == key {
			return res
		}
		res = res.Next
	}
	return nil
}

func (c *LruCache) reorder(node *Node) {
	if node == c.Head {
		return
	}
	node.Prev.Next = node.Next
	if node.Prev == c.Head {
		node.Prev.Prev = node
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	if node.Next == nil {
		c.Tail = node.Prev
	}
	node.Prev = nil
	node.Next = c.Head
	c.Head = node
}

func (c *LruCache) trimPut(key string, value string) {
	node := Node{
		Key:   key,
		Value: value,
	}

	node.Next = c.Head
	c.Head = &node
	c.size++

	if node.Next != nil {
		node.Next.Prev = &node
	}

	if c.size > c.capacity && c.Tail != nil {
		c.Tail = c.Tail.Prev
		if c.Tail != nil {
			c.Tail.Next = nil
		}
		c.size--
	}
	if c.Tail == nil {
		c.Tail = &node
	}
}
