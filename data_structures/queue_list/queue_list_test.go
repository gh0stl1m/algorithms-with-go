package queuelist

import "testing"

func TestStackList(t *testing.T) {
	// Arrange
	var testQueueList *QueueArray = new(QueueArray)
	testQueueList = testQueueList.Init(2)

	// Act
	testQueueList.Queue("Hello")
	testQueueList.Queue("World")

	// Asserts

	if value := testQueueList.Dequeue(); value != "Hello" {
		t.Error("The value is incorrect")
	}

	if value := testQueueList.Dequeue(); value != "World" {
		t.Error("The value is incorrect")
	}

	if value := testQueueList.Dequeue(); value != nil {
		t.Error("The stack must be empty")
	}
}
