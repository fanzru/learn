package main

import (
	"fmt"
)

// Node represents a node in the linked list
type Node struct {
	Val  string
	Next *Node
}

// LinkedList represents the linked list
type LinkedList struct {
	Head *Node
}

// InsertNode inserts a new node with the given value at the beginning of the linked list
func (ll *LinkedList) InsertNode(val string) {
	newNode := &Node{Val: val}
	if ll.Head == nil {
		ll.Head = newNode
	} else {
		newNode.Next = ll.Head
		ll.Head = newNode
	}
}

// PrintList prints all elements in the linked list
func (ll *LinkedList) PrintList() {
	curr := ll.Head
	for curr != nil {
		fmt.Printf("%s -> ", curr.Val)
		curr = curr.Next
	}
	fmt.Println("nil")
}

// TraverseList traverses the linked list and prints each node's value
func (ll *LinkedList) TraverseList() {
	curr := ll.Head
	for curr != nil {
		fmt.Println("Current Node Value:", curr.Val)
		curr = curr.Next
	}
}

func main() {
	// Create a new linked list
	ll := LinkedList{}

	// Insert nodes into the linked list
	ll.InsertNode("Val In")
	ll.InsertNode("Node 2")
	ll.InsertNode("Node 1")

	// Print the linked list
	fmt.Println("Linked List:")
	ll.PrintList()

	// Traverse and print each node's value
	fmt.Println("Traversing the Linked List:")
	ll.TraverseList()
}
