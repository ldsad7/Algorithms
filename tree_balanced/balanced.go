package balanced

import (
	"sort"
)

// A Tree is a binary tree with integer values.
type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func removeDuplicates(elements []int) []int {
	j := 1
	for i := 1; i < len(elements); i++ {
		if elements[i] != elements[i-1] {
			elements[j] = elements[i]
			j++
		}
	}
	return elements[0:j]
}

func RecursiveFunc(elements []int, start int, end int) *Tree {
	if end <= start {
		return nil
	}
	root := new(Tree)
	if end-start > 1 {
		root.Value = elements[start+(end-start)/2-1]
		root.Left = RecursiveFunc(elements, start, start+(end-start)/2-1)
		root.Right = RecursiveFunc(elements, start+(end-start)/2, end)
	} else {
		root.Value = elements[start+(end-start)/2]
		root.Left = RecursiveFunc(elements, start, start+(end-start)/2)
		root.Right = RecursiveFunc(elements, start+(end-start)/2+1, end)
	}
	return root
}

// New constructs a balanced BST containing elements from the given array
func New(elements []int) *Tree {
	if elements == nil || len(elements) == 0 {
		return nil
	}
	sort.Ints(elements)
	elements = removeDuplicates(elements)
	root := &Tree{Value: elements[len(elements)/2]}
	root.Left = RecursiveFunc(elements, 0, len(elements)/2)
	root.Right = RecursiveFunc(elements, len(elements)/2+1, len(elements))
	return root
}
