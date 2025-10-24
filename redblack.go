package redblack

import "errors"

type Node struct {
	Value  int
	Color  string
	Left   *Node
	Right  *Node
	Parent *Node
}

type RedBlackTree struct {
	Root *Node
}

func (t RedBlackTree) Insert(node *Node) error {
	// If node is equal, cannot insert
	// Check if I can continue right or left. If I cannot continue insert on
	// the correct side of currNode
	currNode := t.Root
	for currNode != nil {
		if node.Value == currNode.Value {
			return errors.New("cannot insert duplicate value into tree")
		}
		if node.Value > currNode.Value {
			// Check right
			rightNode := currNode.Right
			if rightNode == nil {
				currNode.Right = node
				if currNode.Parent.Color == "Red" {
					// check for errors
					t.validateTree()
				}
				return nil
			} else {
				currNode = rightNode
			}
		} else {
			leftNode := currNode.Left
			if leftNode == nil {
				currNode.Left = node
				if currNode.Parent.Color == "Red" {
					// check for errors
					t.validateTree()
				}
				return nil
			} else {
				currNode = leftNode
			}
		}
	}
	return nil
}
