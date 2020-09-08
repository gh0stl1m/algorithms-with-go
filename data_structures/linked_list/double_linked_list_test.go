package linkedlist

import "testing"

func TestDoubleLinkedList(t *testing.T) {
	// Arrange
	testList := new(DoubleLinkedList)
	testList.Add(1)
	testList.Add("test")
	testList.Add(3)

	// Act
	testList.Display()
	testList.Remove()
	testList.Display()
}
