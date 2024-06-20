package cache

import (
	"fmt"
	"testing"
)

const (
	storyPageSize = 10 // cache max length
)

func TestCache(t *testing.T) {
	c := NewCache(storyPageSize)
	for i := range storyPageSize * 2 {
		value := i
		c.Add(&Node{Value: value})
	}
	fmt.Println(c)
	t.Log(c)
}
