package cycle

// Node represents an element of singly-linked list
type Node struct {
	Value int
	Next  *Node
}

// HasCycle detects whether the linked list has a cycle
func HasCycle(node *Node) bool {
	if node == nil || node.Next == nil {
		return false
	}
	a := node
	b := node.Next
	for a != nil && b != nil {
		if a == b {
			return true
		}
		a = a.Next
		b = b.Next
		if a == nil || b == nil {
			return false
		}
		b = b.Next
	}
	return false
}
