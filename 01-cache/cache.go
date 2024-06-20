package cache

import (
	"fmt"
	"math"
	"time"
)

// Linked list (Queue) with max size
type cache struct {
	Length, maxLength int
	head, tail        *Node
}

func NewCache(length int) *cache {
	return &cache{maxLength: length}
}

func (c *cache) String() (output string) {
	current := c.GetFirst()
	for i := 0; current != nil; i++ {
		nextValue := math.MinInt
		if current.GetNext() != nil {
			nextValue = current.GetNext().Value
		}
		output = output + fmt.Sprintf("%d: next-> %d | ", current.Value, nextValue)

		current = current.GetNext()
	}
	return output
}

func (c *cache) GetFirst() *Node {
	if c.head == nil {
		return nil
	}
	return c.head
}

func (c *cache) GetLast() *Node {
	if c.tail == nil {
		return nil
	}
	return c.tail
}

// FIFO - first-In-First-Out
// Add to the tail (end) of cache. It'll pop (remove) and return the oldest element if maximum length exceeded.
func (c *cache) Add(element *Node) (popedElement *Node) {
	element.AddedAt = time.Now().UTC()
	if c.head == nil {
		c.head = element
	}
	if c.tail != nil {
		previous := c.tail
		previous.next = element
		element.previous = previous
		c.tail.next = element
	}

	c.tail = element
	c.tail.next = nil

	// pop the oldest element if maximum length exceeded
	if c.Length == c.maxLength {
		popedElement = c.LPop()
	}
	c.Length++
	return popedElement
}

func (c *cache) LPop() *Node {
	fmt.Println(c.head.Value, c.tail.Value, c.Length)
	if c.head == nil {
		return nil
	}
	// check the tail
	// if no more elements remains in the cache
	if c.tail == c.head {
		c.tail = nil
	}
	popElement := c.head
	c.head = c.head.next
	// Set the next of popped element to nil
	popElement.next = nil
	c.Length--
	return popElement
}

func (c *cache) RPop() *Node {
	if c.tail == nil {
		return nil
	}
	// check the head
	// if no more elements remains in the cache
	if c.tail == c.head {
		c.head = nil
	}
	popElement := c.tail
	c.tail = c.tail.previous
	// Set the previous of popped element to nil
	popElement.previous = nil
	c.Length--
	return popElement
}
