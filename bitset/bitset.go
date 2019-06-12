package bitset

import (
	"fmt"
	"math"
)

// Bitset is fixed-size sequence of `size` bits
type Bitset struct {
	size     int
	positive int
	arr      []byte
	flip     int
}

// New creates a new Bitset of a given size
func New(size int) *Bitset {
	return &Bitset{size, 0, make([]byte, size/8), 0}
}

// Set sets a specific bit
func (b *Bitset) Set(pos int, value bool) error {
	if pos >= (*b).size || pos < 0 {
		return fmt.Errorf("Index Error")
	}
	if value == true {
		if (*b).flip == 0 && (*b).arr[pos/8]&byte(math.Pow(2, float64(pos%8))) == 0 {
			(*b).arr[pos/8] |= byte(math.Pow(2, float64(pos%8)))
            (*b).positive++
		} else if (*b).arr[pos/8]&byte(math.Pow(2, float64(pos%8))) != 0 {
			(*b).arr[pos/8] &= byte(255 - math.Pow(2, float64(pos%8)))
            (*b).positive--
		}
	} else {
		if (*b).flip == 0 && (*b).arr[pos/8]&byte(math.Pow(2, float64(pos%8))) != 0 {
			(*b).arr[pos/8] &= byte(255 - math.Pow(2, float64(pos%8)))
            (*b).positive--
		} else if (*b).arr[pos/8]&byte(math.Pow(2, float64(pos%8))) == 0 {
			(*b).arr[pos/8] |= byte(math.Pow(2, float64(pos%8)))
            (*b).positive++
		}
	}
	return nil
}

// Test returns a value of specific bit
func (b *Bitset) Test(pos int) (bool, error) {
	if pos >= (*b).size || pos < 0 {
		return false, fmt.Errorf("Test Error")
	}
	if (*b).flip == 0 {
		return (*b).arr[pos/8]&byte(math.Pow(2, float64(pos%8))) != 0, nil
	}
	return (*b).arr[pos/8]&byte(math.Pow(2, float64(pos%8))) == 0, nil
}

// Count is returns the number of bits set to `true`
func (b *Bitset) Count() int {
	return (*b).positive
}

// All checks if all bits are set to `true`
func (b *Bitset) All() bool {
	return (*b).size == (*b).positive
}

// Any checks if there is at least one bit set to `true`
func (b *Bitset) Any() bool {
	return (*b).positive != 0
}

// Flip toggles the values of bits
func (b *Bitset) Flip() {
	(*b).flip = 1 - (*b).flip
	(*b).positive = (*b).size - (*b).positive
}

// Reset sets all bits to `false`
func (b *Bitset) Reset() {
	for i := 0; i < (*b).size/8; i++ {
		(*b).arr[i] &= 0
	}
	(*b).positive = 0
	(*b).flip = 0
}
