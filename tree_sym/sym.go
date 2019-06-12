package sym

// A Tree is a binary tree with integer values.
type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func IsSymTwoTrees(left *Tree, right *Tree) bool {
	if left == nil || right == nil {
		if right == nil && left == nil {
			return true
		}
		return false
	}
	if left.Value == right.Value {
		return IsSymTwoTrees(left.Left, right.Right) && IsSymTwoTrees(left.Right, right.Left)
	}
	return false
}

// IsSymTree tests whether the given tree is symmetrical
func IsSymTree(tree *Tree) bool {
	if tree == nil {
		return true
	}
	return IsSymTwoTrees(tree.Left, tree.Right)
}
