package trees

import (
	"testing"
)

func TestBST(t *testing.T) {
	// Arrange
	var testBST BST
	testBST.Insert(2)
	testBST.Insert(4)
	testBST.Insert(6)
	testBST.Insert(8)
	testBST.Insert(10)

	// Act
	valueFound := testBST.Find(4)

	// Asserts
	if valueFound != nil {
		t.Error("The value must be found")
	}
}

func TestFindStepsBetween2Nodes(t *testing.T) {
	// Arrange
	var testBST BST
	testBST.Insert(8)
	testBST.Insert(5)
	testBST.Insert(10)
	testBST.Insert(9)
	testBST.Insert(11)
	testBST.Insert(5)
	testBST.Insert(4)
	testBST.Insert(6)

	// Act
	valueFound := testBST.FindSteps(4, 6)

	// Asserts
	if valueFound != 1 {
		t.Error("The number steps is invalid")
	}
}

func TestIfTwoTreesAreEqual(t *testing.T) {
	// Arrange
	var testBST1 BST
	testBST1.Insert(4)
	testBST1.Insert(3)
	testBST1.Insert(6)

	var testBST2 BST
	testBST2.Insert(6)
	testBST2.Insert(3)
	testBST2.Insert(4)

	// Act
	areTreesEqual := AreEqual(&testBST1, &testBST2)

	// Asserts
	if areTreesEqual != true {
		t.Error("Trees must be equal")
	}
}

func TestPrintByLevel(t *testing.T) {
	// Arrange
	var testBST BST
	testBST.Insert(8)
	testBST.Insert(5)
	testBST.Insert(10)
	testBST.Insert(9)
	testBST.Insert(11)
	testBST.Insert(5)
	testBST.Insert(4)
	testBST.Insert(6)

	// Act
	testBST.PrintByLevel(3)
}
