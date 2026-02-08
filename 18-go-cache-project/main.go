package main

import (
	"fmt"
)

const size = 5

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
	Hash  Hashs
}

type Hashs map[string]*Node

func NewCache() *Cache {
	return &Cache{Queue: NewQueue(), Hash: Hashs{}}
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
	tmp := c.Queue.Head.right

	c.Queue.Head.right = n
	n.left = c.Queue.Head
	n.right = tmp
	n.left = n

	c.Queue.Length++

	if c.Queue.Length > size {
		c.Remove(c.Queue.Tail.left)
	}
}

func (c *Cache) Display() {
	c.Queue.Display()

}

func (q *Queue) Display() {
	node := q.Head.right
	fmt.Printf("%d - [", q.Length)
	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", node.Val)
		if i < q.Length-1 {
			fmt.Printf("<-->")
		}
		node = node.right
	}
	fmt.Println("]")
}

func main() {
	fmt.Println("Go Cache Project")
	cache := NewCache()
	for _, words := range []string{"capsicum", "carror", "guva", "tomato", "capsicum", "warrier", "universe"} {
		cache.Check(words)
		cache.Display()
	}
}
