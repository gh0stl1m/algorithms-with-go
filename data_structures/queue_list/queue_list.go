package queuelist

import "fmt"

// QueueArray holds the struck of the queue
type QueueArray struct {
	ArrayQueue []interface{}
	Front      int
	Rear       int
	Size       uint
}

// Init create a instance of the queue array
func (qa *QueueArray) Init(size uint) *QueueArray {
	return &QueueArray{
		ArrayQueue: make([]interface{}, size),
		Front:      -1,
		Rear:       -1,
		Size:       size,
	}
}

// Queue adds an element into the queue
func (qa *QueueArray) Queue(item interface{}) {
	if qa.isFull() {
		fmt.Println("The queue is full")
	}
	if qa.Front == -1 {
		qa.Front = 0
	}
	qa.Rear++
	qa.ArrayQueue[qa.Rear] = item
}

// Dequeue removes an element into the queue
func (qa *QueueArray) Dequeue() interface{} {
	if qa.isEmpty() {
		fmt.Println("The queue is empty")
		return nil
	}

	item := qa.ArrayQueue[qa.Front]
	qa.Front++

	return item
}

func (qa *QueueArray) isFull() bool {
	return qa.Rear == int(qa.Size-1)
}

func (qa *QueueArray) isEmpty() bool {
	return qa.Front == -1 || qa.Front > qa.Rear
}

// Problem: Friend circle

// FindFriendsCircle retrieves the friends circles given a matrix of connection
func FindFriendsCircle(connections [][]uint) [][]int {
	var listOfCircles [][]int
	var visitedConn []bool = make([]bool, len(connections))

	for pos := range connections {
		if visitedConn[pos] == true {
			continue
		}

		visitedConn[pos] = true

		var listOfFriends []int
		var circleQueue *QueueArray = new(QueueArray)
		circleQueue = circleQueue.Init(uint(len(connections)))
		circleQueue.Queue(pos)

		for !circleQueue.isEmpty() {
			userID := circleQueue.Dequeue()
			listOfFriends = append(listOfFriends, userID.(int))

			for friendID := range connections[pos] {
				if connections[userID.(int)][friendID] == 1 && !visitedConn[friendID] {
					circleQueue.Queue(friendID)
					visitedConn[friendID] = true
				}
			}
		}

		listOfCircles = append(listOfCircles, listOfFriends)
	}

	return listOfCircles
}
