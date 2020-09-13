package trees

import "fmt"

// Node define the base structure
type Node struct {
	Value int
	Right *Node
	Left  *Node
}

// BST defines the structure of the binary search tree
type BST struct {
	Root *Node
}

// Insert adds the root node of the BST
func (bst *BST) Insert(value int) {
	if bst.Root == nil {
		bst.Root = &Node{
			Value: value,
		}
	} else {
		bst.Root.Insert(value)
	}
}

// Insert adds a value into the tree
func (n *Node) Insert(value int) {
	if value <= n.Value {
		if n.Left == nil {
			n.Left = &Node{
				Value: value,
			}
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

// Find retrieves an element inside a BST
func (bst *BST) Find(value int) error {
	node := bst.findNode(bst.Root, value)

	if node == nil {
		return fmt.Errorf("Value: %d not found in tree", value)
	}

	return nil
}

func (bst *BST) findNode(node *Node, value int) *Node {
	if node == nil {
		fmt.Println("Node not found")
		return nil
	}

	if node.Value == value {
		fmt.Println("Node is root")
		return bst.Root
	}

	fmt.Println("Node value: ", node.Value)

	if value < node.Value {
		return bst.findNode(node.Left, value)
	}

	return bst.findNode(node.Right, value)
}
