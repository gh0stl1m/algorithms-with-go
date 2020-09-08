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

func TestDeleteKnElement(t *testing.T) {
	// Arrange
	testList := new(SimpleLinkedList)
	testList.Add(1)
	testList.Add(2)
	testList.Add(3)
	testList.Add(4)
	testList.Add(5)
	testList.Add(6)

	// Act
	DeleteKnElement(testList, 2)
	testList.Display()
}
