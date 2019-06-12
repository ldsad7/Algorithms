package array

import (
	"fmt"
	"strconv"
)

// Element is a type of an array element
type Element uint64

// Array is an implementation of list using expandable array
// with fast insertion to and deletion from the end
type Array struct {
	elems []Element
	n     int
	cap   int
}

// New creates a new Array with a given capacity
func New(cap int) *Array {
	elems := make([]Element, cap, cap)
	return &Array{elems, 0, cap}
}

// Len returns the length of the array
func (a *Array) Len() int {
	return a.n
}

// Get retrieves an array element by index
func (a *Array) Get(i int) (Element, error) {
	var elem Element

	if i >= 0 && i < a.n {
		return a.elems[i], nil
	}
	return elem, fmt.Errorf("Get Error")
}

// Set writes an element to array by index
func (a *Array) Set(i int, x Element) error {
	if i >= 0 && i < a.n {
		a.elems[i] = x
		return nil
	}
	return fmt.Errorf("Set Error")
}

// Insert adds an element to the array by index
func (a *Array) Insert(i int, x Element) error {
	if i >= a.n+1 || i < 0 {
		return fmt.Errorf("Insert Error")
	}
	q := 2
	if a.n == a.cap || len(a.elems) == i {
		new := make([]Element, a.n*q+1, a.n*q+1)
		copy(new, a.elems)
		a.elems = nil
		a.elems = make([]Element, a.n*q+1, a.n*q+1)
		copy(a.elems, new)
		new = nil
		a.cap = a.n*q + 1
	}
	a.n++
	for ; i < a.n; i++ {
		tmp := a.elems[i]
		a.elems[i] = x
		x = tmp
	}
	return nil
}

// Push inserts an element to the right end of the array
func (a *Array) Push(x Element) error {
	err := a.Insert(a.Len(), x)
	if err != nil {
		return err
	}
	return nil
}

// Delete removes an element from the array by index
func (a *Array) Delete(i int) error {
	j := a.Len() - 1
	if j < 0 || i < 0 || i > j {
		return fmt.Errorf("Delete Error")
	}
	for x := a.elems[j]; j > i; j-- {
		tmp := a.elems[j-1]
		a.elems[j-1] = x
		x = tmp
	}
	a.n--
	if a.n == 0 {
		a.elems = nil
	}
	return nil
}

// Pop deletes the last element of the array
func (a *Array) Pop() error {
	err := a.Delete(a.Len() - 1)
	return err
}

// String returns a textual representation of the array
func (a *Array) String() string {
	s := "["
	i := 0
	for ; i < a.Len()-1; i++ {
		tmp, _ := a.Get(i)
		s += strconv.Itoa(int(tmp))
		s += ", "
	}
	tmp, err := a.Get(i)
	if err == nil {
		s += strconv.Itoa(int(tmp))
	}
	s += "]"
	return s
}
