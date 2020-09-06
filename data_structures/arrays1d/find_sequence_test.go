package arrays1d

import (
	"testing"
)

func TestFindLongestSequence(t *testing.T) {
	// Arrange
	arrayInput1 := []uint{1, 1, 0, 1, 1, 1, 1, 0, 1, 1}
	arrayInput2 := []uint{1, 1, 0, 1, 0, 1, 1, 0, 1, 1}
	var expectedLongestSequence1 uint = 7
	var expectedLongestSequence2 uint = 5
	// expectedArr := []uint{2, 0, 4, 0, 2}

	// Act
	longestSequence1 := FindLongestSequence(arrayInput1)
	longestSequence2 := FindLongestSequence(arrayInput2)
	// arraysAreEqual := reflect.DeepEqual(valuesMapped, expectedArr)

	// Asserts
	if longestSequence1 != expectedLongestSequence1 {
		t.Error("The longest squence values is wroing")
	}

	if longestSequence2 != expectedLongestSequence2 {
		t.Error("The longest squence values is wroing")
	}
}
