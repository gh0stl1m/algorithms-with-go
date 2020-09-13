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
