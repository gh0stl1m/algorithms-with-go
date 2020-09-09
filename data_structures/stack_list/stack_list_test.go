package stacklist

import "testing"

func TestStackList(t *testing.T) {
	// Arrange
	var testStackList *StackArray = new(StackArray)
	testStackList = testStackList.Init(2)

	// Act
	testStackList.Push("Hello")
	testStackList.Push("World")

	// Asserts

	if value := testStackList.Pop(); value != "World" {
		t.Error("The value is incorrect")
	}

	if value := testStackList.Pop(); value != "Hello" {
		t.Error("The value is incorrect")
	}

	if value := testStackList.Pop(); value != nil {
		t.Error("The stack must be empty")
	}
}
