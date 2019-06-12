package reverse

// Node represents an element of singly-linked list
type Node struct {
	Value int
	Next  *Node
}

// Reverse flips the original order of elements in the singly-linked list and returns a pointer to the resulting list
func Reverse(node *Node) *Node {
	var a *Node
	var b *Node = nil
	for node != nil && node.Next != nil {
		a = node.Next
		node.Next = b
		b = node
		node = a
	}
	if node != nil {
		node.Next = b
	}
	return node
}
