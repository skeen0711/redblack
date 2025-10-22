package redblack

import "testing"

func TestNode(t *testing.T) {
	root := &Node{1, "Black", nil, nil, nil}
	root.Left = &Node{2, "Red", nil, nil, root}
	root.Right = &Node{3, "Red", nil, nil, root}

	rightColor := root.Right.Color
	rightValue := root.Right.Value
	if rightColor != "Red" {
		t.Errorf("Expected right node to be true (red), got %v", rightColor)
	}
	if rightValue != 3 {
		t.Errorf("Expected right node to be 3, got %v", rightValue)
	}
}

func TestInsertRedBlackTree(t *testing.T) {
	root := &Node{1, "Black", nil, nil, nil}
	tree := RedBlackTree{root}
	left := &Node{0, "Red", nil, nil, root}
	right := &Node{2, "Red", nil, nil, root}
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
	if newRoot.Left.Color != "Red" {
		t.Errorf("Expected left node to be true (red), got %v", newRoot.Left.Color)
	}
	if newRoot.Right.Color != "Red" {
		t.Errorf("Expected right node to be true (red), got %v", newRoot.Right.Color)
	}
}
