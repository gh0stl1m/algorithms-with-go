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

// FindWordsPath retrieve true if there's a path between strtd and endword
func FindWordsPath(dictionary []string, startWord string, endWord string, wordDistance uint) bool {
	var isPathFound bool
	var visitedDic []bool = make([]bool, len(dictionary))
	var stackPath *StackArray = new(StackArray)
	stackPath = stackPath.Init(uint(len(dictionary)))
	stackPath.Push(startWord)

	for !stackPath.isEmpty() {
		lastWord := stackPath.Pop()

		if getWordDistance(wordDistance, startWord, endWord) == 0 {
			stackPath.Push(lastWord)
			isPathFound = true
			break
		}

		for pos := range dictionary {
			if visitedDic[pos] == true {
				continue
			}

			distance := getWordDistance(wordDistance, lastWord.(string), dictionary[pos])
			if distance == 1 {
				visitedDic[pos] = true
				stackPath.Push(lastWord)
				stackPath.Push(dictionary[pos])
			}

		}
	}

	return isPathFound
}

func getWordDistance(distance uint, word1 string, word2 string) uint {
	var wordDistance uint = distance
	for pos := range word1 {
		if word1[pos] == word2[pos] {
			wordDistance--
		}
	}

	return wordDistance
}

// IsValidParenthesesExpresions validates is a parenthesis expression is valid
func IsValidParenthesesExpresions(expresion []string) bool {
	var stack1 *StackArray = new(StackArray)
	stack1 = stack1.Init(uint(len(expresion)))
	var stack2 *StackArray = new(StackArray)
	stack2 = stack2.Init(uint(len(expresion)))
	var isValid bool = true

	for _, value := range expresion {
		stack1.Push(value)
	}

	for !stack1.isEmpty() {
		parenthesis := stack1.Pop()

		if parenthesis == ")" {
			stack2.Push(parenthesis)
		} else {
			if stack2.isEmpty() {
				isValid = false
				break
			} else {
				stack2.Pop()
			}
		}
	}

	return isValid
}
