package queuelist

import (
	"reflect"
	"testing"
)

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

func TestFriendsCircles(t *testing.T) {
	// Arrange
	var connections [][]uint = [][]uint{{1, 1, 0, 0, 0}, {1, 1, 1, 0, 0}, {0, 1, 1, 0, 0}, {0, 1, 1, 0, 0}}
	var expectedConnections [][]uint = [][]uint{{0, 1, 2}, {3, 4}}

	// Act
	listOfFriends := FindFriendsCircle(connections)
	listIsEqual := reflect.DeepEqual(listOfFriends, expectedConnections)

	// Asserts
	if listIsEqual {
		t.Error("The list of friend must be equal")
	}

}
