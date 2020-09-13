package trees

import (
	"fmt"
)

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

// Problem: Find the number of steps between two nodes

// FindSteps retrieves the number of steps that are between two nodes
func (bst *BST) FindSteps(nodeValue1, nodeValue2 int) int {
	var steps int = 0
	var nodesVisited []int

	findPath(nodeValue1, &steps, &nodesVisited, bst.Root)
	findPath(nodeValue2, &steps, &nodesVisited, bst.Root)

	fmt.Println("Nodes visited: ", nodesVisited)

	return steps
}

// Problem: Check if two trees are equal

// AreEqual check if two trees are equal
func AreEqual(tree1, tree2 *BST) bool {
	var areTreesEqual bool = true
	var nodesMap map[int]int = make(map[int]int)

	treeChecker(tree1.Root, &nodesMap)
	treeChecker(tree2.Root, &nodesMap)

	for _, value := range nodesMap {
		if value != 2 {
			areTreesEqual = false
			break
		}
	}

	return areTreesEqual
}

func treeChecker(currentNode *Node, nodes *map[int]int) {
	if currentNode == nil {
		return
	}
	treeChecker(currentNode.Right, nodes)
	(*nodes)[currentNode.Value]++
	treeChecker(currentNode.Left, nodes)
}

func findPath(searchValue int, steps *int, visitedNodes *[]int, currentNode *Node) {
	if !contains(currentNode.Value, *visitedNodes) {
		*steps++
		*visitedNodes = append(*visitedNodes, currentNode.Value)
	} else {
		*steps--
	}

	if currentNode.Value == searchValue {
		return
	}

	if searchValue > currentNode.Value {
		findPath(searchValue, steps, visitedNodes, currentNode.Right)
	} else {
		findPath(searchValue, steps, visitedNodes, currentNode.Left)
	}
}

func contains(key int, source []int) bool {
	for _, value := range source {
		if key == value {
			return true
		}
	}
	return false
}
