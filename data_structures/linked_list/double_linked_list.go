package linkedlist

import "fmt"

// NodeD represents the pointer to the next value
type NodeD struct {
	Value interface{}
	Next  *NodeD
	Prev  *NodeD
}

// DoubleLinkedList represents a double linked list
type DoubleLinkedList struct {
	Head *NodeD
	Len  uint
}

// Add Creates a new element into the linked list
func (dl *DoubleLinkedList) Add(value interface{}) {
	newNode := &NodeD{
		Value: value,
	}

	if dl.Head != nil {
		newNode.Next = dl.Head
		dl.Head.Prev = newNode
		dl.Head = newNode
	} else {
		dl.Head = newNode
	}

	dl.Len++
}

// Remove eliminate the head element of the list
func (dl *DoubleLinkedList) Remove() error {
	if dl.Head == nil {
		return fmt.Errorf("List is empty")
	}

	dl.Head = dl.Head.Next
	dl.Head.Prev = nil
	dl.Len--

	return nil
}

// Display prints the linked list
func (dl *DoubleLinkedList) Display() {
	node := dl.Head

	for node != nil {
		fmt.Printf("%v -> ", node.Value)
		node = node.Next
	}
	fmt.Println()
}
