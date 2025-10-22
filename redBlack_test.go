package redblack

import "testing"

func TestNode(t *testing.T) {
	root := &Node{1, false, nil, nil}
	root.Left = &Node{2, true, nil, nil}
	root.Right = &Node{3, true, nil, nil}

	rightColor := root.Right.isRed
	rightValue := root.Right.Value
	if rightColor != true {
		t.Errorf("Expected right node to be true (red), got %v", rightColor)
	}
	if rightValue != 3 {
		t.Errorf("Expected right node to be 3, got %v", rightValue)
	}
}
