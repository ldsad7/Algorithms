package inorder

// A Tree is a binary tree with integer values.
type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

// Inorder returns an inorder traversal of the given tree
func Inorder(tree *Tree) []int {
	arr := []int{}
	for tree != nil {
		if tree.Left != nil {
			tmp := tree.Left
			for tmp.Right != tree && tmp.Right != nil {
				tmp = tmp.Right
			}
			if tmp.Right == nil {
				tmp.Right = tree
				tree = tree.Left
			} else {
				arr = append(arr, tree.Value)
				tree = tree.Right
			}
		} else {
			arr = append(arr, tree.Value)
			tree = tree.Right
		}
	}
	return arr
}
