package hashmap

import "testing"

func TestFindUniqueNumber(t *testing.T) {
	// Arrange
	var arrayOfValues1 []int = []int{1, 2, 2, 5, 5, 1, 4}
	var arrayOfValues2 []int = []int{1, 2, 2, 4, 5, 1, 4, 6}

	// Act
	uniqueNumber1 := FindUniqueNumber(arrayOfValues1)
	uniqueNumber2 := FindUniqueNumber(arrayOfValues2)

	// Asserts
	if uniqueNumber1 != 4 {
		t.Error("Unexpected value")
	}

	if uniqueNumber2 != 6 {
		t.Error("Unexpected value")
	}
}

func TestFindRelevantWord(t *testing.T) {
	// Arrange
	var inputText string = "Suppose we have a set of English text document and wish to rank which document is more relevant to the query, the brown cow. A simple way to start out is by eliminating documents that do not contain all three words the brown, and cow, but this still leaves many documents."
	var wordsToExclude []string = []string{"of", "we", "a", "and", "to", "is", "the", ",", "A", "by", "not"}

	// Act
	relevantWord := GetMostRelevantWord(inputText, wordsToExclude)

	// Asserts
	if relevantWord != "document" {
		t.Error("The word is not the most relevant in the input text")
	}
}
