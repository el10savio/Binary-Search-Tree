package main

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (node *Node) Search(Key int) bool {
	if node == nil {
		return false
	}

	if node.Key < Key {
		return node.Right.Search(Key)
	} else {
		return node.Left.Search(Key)
	}

	return true
}

func (node *Node) Insert(Key int) {
	if node.Key < Key {
		if node.Right == nil {
			node.Right = &Node{Key: Key}
		} else {
			node.Right.Insert(Key)
		}
	} else if node.Key > Key {
		if node.Left == nil {
			node.Left = &Node{Key: Key}
		} else {
			node.Left.Insert(Key)
		}
	}
}

func (node *Node) Delete(Key int) *Node {
	if node.Key < Key {
		node.Right = node.Right.Delete(Key)
	} else if node.Key > Key {
		node.Left = node.Left.Delete(Key)
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}
		min := node.Right.Min()
		node.Key = min
		node.Right = node.Right.Delete(min)
	}
	return node
}

func (node *Node) Min() int {
	if node.Left == nil {
		return node.Key
	}
	return node.Left.Min()
}

func (node *Node) Max() int {
	if node.Right == nil {
		return node.Key
	}
	return node.Right.Max()
}

func main() {

}
