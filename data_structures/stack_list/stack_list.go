package stacklist

import "fmt"

// StackArray holds the struck of the stack list
type StackArray struct {
	ArrayStack []interface{}
	Size       uint
	Top        int
}

// Init initialize the stack array
func (sa *StackArray) Init(size uint) *StackArray {
	return &StackArray{
		ArrayStack: make([]interface{}, size),
		Size:       size,
		Top:        -1,
	}
}

// Push adds a new item to the stack list
func (sa *StackArray) Push(newItem interface{}) {
	if sa.isFull() {
		fmt.Println("Stack is full")
		return
	}

	sa.Top++
	sa.ArrayStack[sa.Top] = newItem
}

// Pop retrieves the last item in the stack
func (sa *StackArray) Pop() interface{} {
	if sa.isEmpty() {
		fmt.Println("The stack is empty")
		return nil
	}

	item := sa.ArrayStack[sa.Top]
	sa.Top--

	return item
}

func (sa *StackArray) isFull() bool {
	return uint(sa.Top) == (sa.Size - 1)
}

func (sa *StackArray) isEmpty() bool {
	return sa.Top == -1
}
