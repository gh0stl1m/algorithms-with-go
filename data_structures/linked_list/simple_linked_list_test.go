package linkedlist

import "testing"

func TestSimpleLinkedList(t *testing.T) {
	// Arrange
	testList := new(SimpleLinkedList)
	testList.Add(1)
	testList.Add("test")
	testList.Add(3)

	// Act
	testList.Display()
	testList.Remove()
	testList.Display()
}
