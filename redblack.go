package redblack

type RedBlackTree struct {
	Root *Node
}

type Node struct {
	Key    int
	Value  string
	Color  Color
	Left   *Node
	Right  *Node
	Parent *Node
}
type Color bool

const (
	Black Color = true
	Red   Color = false
)

func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{}
}

func (t *RedBlackTree) Insert(key int, value string) {
	newNode := &Node{Key: key, Value: value, Color: Red}

	if t.Root == nil {
		t.Root = newNode
		newNode.Color = Black // Root is always black
		return
	}

	currNode := t.Root
	for currNode != nil {
		if key < currNode.Key {
			if currNode.Left == nil {
				currNode.Left = newNode
				newNode.Parent = currNode
				break
			}
			currNode = currNode.Left
		} else if key > currNode.Key {
			if currNode.Right == nil {
				currNode.Right = newNode
				newNode.Parent = currNode
				break
			}
			currNode = currNode.Right
		} else {
			currNode.Value = value
			return
		}
	}
}

func (t *RedBlackTree) Search(key int) (string, bool) {
	node := t.Root
	for node != nil {
		if key < node.Key {
			node = node.Left
		} else if key > node.Key {
			node = node.Right
		} else {
			return node.Value, true
		}
	}
	return "", false
}
