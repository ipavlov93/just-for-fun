package cache

import "time"

type Node struct {
	Value          int
	AddedAt        time.Time
	next, previous *Node
}

func (c *Node) GetNext() *Node {
	if c == nil {
		return nil
	}
	return c.next
}
