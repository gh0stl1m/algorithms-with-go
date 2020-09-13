package trees

import "testing"

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

	if valueFound != 1 {
		t.Error("The number steps is invalid")
	}
}
