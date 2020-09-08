package linkedlist

import "fmt"

// Node represents the pointer to the next value
type Node struct {
	Value interface{}
	Next  *Node
}

// SimpleLinkedList represents a simple linked list
type SimpleLinkedList struct {
	Head *Node
	Len  uint
}

// Add Creates a new element into the linked list
func (spl *SimpleLinkedList) Add(value interface{}) {
	newNode := &Node{
		Value: value,
	}

	if spl.Head != nil {
		newNode.Next = spl.Head
		spl.Head = newNode
	} else {
		spl.Head = newNode
	}

	spl.Len++
}

// Remove eliminate the head element of the list
func (spl *SimpleLinkedList) Remove() error {
	if spl.Head == nil {
		return fmt.Errorf("List is empty")
	}

	spl.Head = spl.Head.Next
	spl.Len--
	return nil
}

// Display prints the linked list
func (spl *SimpleLinkedList) Display() {
	node := spl.Head

	for node != nil {
		fmt.Printf("%v -> ", node.Value)
		node = node.Next
	}
	fmt.Println()
}
