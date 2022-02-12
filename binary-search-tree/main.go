package main

import "fmt"

// Node
type Node struct {
	Key int 
	Left *Node
	Right *Node
}

// Insert - add a note to a tree. The key can't already be in the tree
func (n *Node) Insert (k int) {
	if n.Key < k {
		// Move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k{
		// Move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Search - searches for a node with this key in the tree. If it's found, it will return true
func (n *Node) Search (k int) bool {
	if n == nil {
		return false
	}
	
	if n.Key < k {
		// Move Right
		return n.Right.Search(k)
	} else if n.Key > k {
		// Move Left
		return n.Left.Search(k)
	}

	return true
}

func main() {
	tree := &Node{Key: 100}
	fmt.Println(tree)

	tree.Insert(52)
	tree.Insert(303)
	tree.Insert(19)
	tree.Insert(76)
	tree.Insert(310)
	tree.Insert(88)
	tree.Insert(276)
	tree.Insert(7)
	tree.Insert(11)

	fmt.Println(tree.Search(276))
}