package deduplicate

// Node represents an element of singly-linked list
type Node struct {
	Value uint32
	Next  *Node
}

// Deduplicate removes duplicates from a sorted singly-linked list starting with `node`
func Deduplicate(node *Node) {
	for node != nil && node.Next != nil {
		if node.Value == node.Next.Value {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}
}
