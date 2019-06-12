package leftleaves

// A Tree is a binary tree with integer values.
type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func SumOfRight(tree *Tree) int {
	if tree == nil {
		return 0
	}
	return SumOfLeft(tree.Left) + SumOfRight(tree.Right)
}

func SumOfLeft(tree *Tree) int {
	if tree == nil {
		return 0
	} else if tree.Left == nil && tree.Right == nil {
		return tree.Value
	}
	return SumOfLeft(tree.Left) + SumOfRight(tree.Right)
}

//
// SumOfLeftLeaves returns the sum of left leaves values for the given tree
func SumOfLeftLeaves(tree *Tree) int {
	if tree == nil {
		return 0
	}
	return SumOfLeft(tree.Left) + SumOfRight(tree.Right)
}
