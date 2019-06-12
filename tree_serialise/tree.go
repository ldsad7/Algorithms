package tree

import (
	"fmt"
	"strconv"
)

// A Tree is a binary tree with integer values.
type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func New_Tree(data []string, length int, i int, tree *Tree) (*Tree, error) {
	if 2*i+1 < length {
		if tree == nil && data[2*i+1] != "nil" {
			return nil, fmt.Errorf("My Error")
		}
		var tmp int
		var err error
		if data[2*i+1] != "nil" {
			tmp, err = strconv.Atoi(data[2*i+1])
			if err == nil {
				left := Tree{Value: tmp}
				tree.Left, err = New_Tree(data, length, 2*i+1, &left)
				if err != nil {
					return nil, err
				}
			} else {
				return nil, err
			}
		} else {
			_, err = New_Tree(data, length, 2*i+1, nil)
			if err != nil {
				return nil, err
			}
		}
	}
	if 2*i+2 < length {
		if tree == nil && data[2*i+2] != "nil" {
			return nil, fmt.Errorf("My Error")
		}
		var tmp int
		var err error
		if data[2*i+2] != "nil" {
			tmp, err = strconv.Atoi(data[2*i+2])
			if err == nil {
				right := Tree{Value: tmp}
				tree.Right, err = New_Tree(data, length, 2*i+2, &right)
				if err != nil {
					return nil, err
				}
			} else {
				return nil, err
			}
		} else {
			_, err = New_Tree(data, length, 2*i+2, nil)
			if err != nil {
				return nil, err
			}
		}
	}
	return tree, nil
}

// New creates a binary tree from the array representation
func New(data []string) (*Tree, error) {
	if data == nil || len(data) == 0 {
		return nil, nil
	}
	length := len(data)
	var tree *Tree
	if data[0] != "nil" {
		tmp, err := strconv.Atoi(data[0])
		if err != nil {
			return nil, err
		}
		tree := Tree{Value: tmp}
		return New_Tree(data, length, 0, &tree)
	} else {
		tree = nil
		return New_Tree(data, length, 0, tree)
	}
}

func Serialise_Tree(data []string, i int, tree *Tree) []string {
	length := len(data)
	if tree.Left != nil {
		for length < 2*i+2 {
			data = append(data, "")
			length++
		}
		data[2*i+1] = strconv.Itoa(tree.Left.Value)
		data = Serialise_Tree(data, 2*i+1, tree.Left)
	}
	length = len(data)
	if tree.Right != nil {
		for length < 2*i+3 {
			data = append(data, "")
			length++
		}
		data[2*i+2] = strconv.Itoa(tree.Right.Value)
		data = Serialise_Tree(data, 2*i+2, tree.Right)
	}
	return data
}

// Serialise returns a normalised array representation of the given tree
func Serialise(tree *Tree) []string {
	str := []string{}
	if tree == nil {
		return str
	}
	str = append(str, strconv.Itoa(tree.Value))
	str = Serialise_Tree(str, 0, tree)
	for i := range str {
		if str[i] == "" {
			str[i] = "nil"
		}
	}
	return str
}
