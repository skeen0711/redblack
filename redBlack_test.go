package redblack

import "testing"

func TestNode(t *testing.T) {
	root := &Node{1, false, nil, nil, nil}
	root.Left = &Node{2, true, nil, nil, root}
	root.Right = &Node{3, true, nil, nil, root}

	rightColor := root.Right.isRed
	rightValue := root.Right.Value
	if rightColor != true {
		t.Errorf("Expected right node to be true (red), got %v", rightColor)
	}
	if rightValue != 3 {
		t.Errorf("Expected right node to be 3, got %v", rightValue)
	}
}

func TestInsertRedBlackTree(t *testing.T) {
	root := &Node{1, false, nil, nil, nil}
	tree := RedBlackTree{root}
	left := &Node{0, true, nil, nil, root}
	right := &Node{2, true, nil, nil, root}
	rerr := tree.Insert(left)
	if rerr != nil {
		return
	}
	lerr := tree.Insert(right)
	if lerr != nil {
		return
	}
	newRoot := tree.Root
	if newRoot.Left.Value != 0 {
		t.Errorf("Expected left node to be 0, got %v", newRoot.Left.Value)
	}
	if newRoot.Right.Value != 2 {
		t.Errorf("Expected right node to be 2, got %v", newRoot.Right.Value)
	}
	if newRoot.Left.isRed != true {
		t.Errorf("Expected left node to be true (red), got %v", newRoot.Left.isRed)
	}
	if newRoot.Right.isRed != true {
		t.Errorf("Expected right node to be true (red), got %v", newRoot.Right.isRed)
	}
}
