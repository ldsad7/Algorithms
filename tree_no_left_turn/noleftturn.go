package noleftturn

// A Tree is a binary tree with integer values.
type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

// NoLeftTurn transforms a binary search tree into an equivalent one that has no left subtrees.
// Returns a pointer to the new root.
func NoLeftTurn(tree *Tree) *Tree {
	if tree == nil || (tree.Left == nil && tree.Right == nil) {
		return tree
	} else if tree.Left == nil {
		tree.Right = NoLeftTurn(tree.Right)
		return tree
	} else {
		tree.Right = NoLeftTurn(tree.Right)
		top := NoLeftTurn(tree.Left)
		tmp := top
		for tmp.Right != nil {
			tmp = tmp.Right
		}
		tmp.Right = tree
		tree.Left = nil
		return top
	}
}
