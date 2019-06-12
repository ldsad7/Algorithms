package ndarray

import (
	"fmt"
)

// NDArray is a prototype of an n-dimensional array
type NDArray struct {
	shape []uint32
}

// New creates an NDArray with given dimensions
func New(shape ...uint32) *NDArray {
	ndarr := new(NDArray)
	ndarr.shape = shape
	return ndarr
}

// Idx an index in a linear array for a given n-dimensional index
func (nda *NDArray) Idx(indices ...uint32) (uint32, error) {
	if len(nda.shape) != len(indices) {
		return 0, fmt.Errorf("Error")
	}
	var sum uint32 = 0
	var tmp uint32 = 1
	for i := len(indices) - 1; i >= 0; i-- {
		if indices[i] >= nda.shape[i] {
			return 0, fmt.Errorf("Error")
		}
		sum += indices[i] * tmp
		tmp *= nda.shape[i]
	}
	return sum, nil
}
