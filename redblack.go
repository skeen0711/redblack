package redblack

type Node struct {
	Value int
	isRed bool
	Left  *Node
	Right *Node
}

type RedBlackTree struct {
	Root *Node
}
