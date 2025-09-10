package main

import (
	"fmt"
)

type Node struct {
	Val   string
	left  *Node
	right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Cache struct {
	Queue Queue
	Hash  Hash
}

type Hash map[string]*Node

func NewCache() *Cache {
	return Cache{Queue: NewQueue, Hash: Hash{}}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.right = tail
	tail.left = head

	return Queue{Head: head, Tail: tail}
}

func (c *Cache) Check(word string) {
	node := &Node{}

	if val, ok := c.Hash[word]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Val: word}
	}

	c.Add(node)
	c.Hash[word] = node
}

func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("Removing %s\n", n.Val)
	left := n.left
	right := n.right

	left.right = right
	right.left = left
	c.Queue.Length--
	delete(c.Hash, n.Val)
	return n
}

func (c *Cache) Add(n *Node) {
	fmt.Printf("add: %s\n", n.Val)
	tmp := c.Queue.Head.Right

	c.Queue.Head.Right = n
}

func main() {
	fmt.Println("Go Cache Project")
	cache := NewCache()
	for _, words := range []string{"Hello", "carror", "guva", "tomato", "capsicum"} {
		cache.Check(words)
		cache.Display()
	}
}
