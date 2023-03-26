package binarysearchtree

type Node struct {
	Value int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func (b *BST) Insert(data int) {
	node := Node{Value: data}
	if b.root == nil {
		b.root = &node
		return
	}
	current := b.root
	for b.root != nil {
		if data < current.Value {
			// move left
			if current.left == nil {
				current.left = &node
				return
			} else {
				current = current.left
			}
		} else {
			// move right
			if current.right == nil {
				current.right = &node
				return
			} else {
				current = current.right
			}
		}
	}
}

func (b *BST) Lookup(data int) *Node {
	if b.root == nil {
		return nil
	}
	current := b.root
	for current != nil {
		if data < current.Value {
			current = current.left
		} else if data > current.Value {
			current = current.right
		} else if data == current.Value {
			return current
		}
	}
	return nil
}