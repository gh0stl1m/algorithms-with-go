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

// Problem: Word distance
func TestWordDistance(t *testing.T) {
	// Arrange
	var dictionary []string = []string{"hot", "dot", "dog", "lot", "log", "cog"}
	var startWord string = "hot"
	var endWord string = "cog"

	// Act
	isPathFound := FindWordsPath(dictionary, startWord, endWord, 3)

	// Asserts
	if isPathFound {
		t.Error("The path must be found")
	}
}

func TestIsValidParethesesExpression(t *testing.T) {
	// Arrange
	var expression1 []string = []string{"(", "(", ")", ")"}
	var expression2 []string = []string{"(", "(", "(", ")", ")"}

	// Act
	isValidExpression := IsValidParenthesesExpresions(expression1)
	isNotValidExpression := IsValidParenthesesExpresions(expression2)

	// Asserts
	if !isValidExpression {
		t.Error("The expression must be valid")
	}

	if isNotValidExpression {
		t.Error("The expression must not be valid")
	}
}
