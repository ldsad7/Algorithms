package streamlargest

import "fmt"

type Stream struct {
	k       int
	len     int
	MinHeap []int
}

// New creates a Stream instance
func New(k int) (*Stream, error) {
	if k > 0 {
		stream := &(Stream{0, k, make([]int, k)})
		return stream, nil
	} else {
		return nil, fmt.Errorf("New Error")
	}
}

func Min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

// Push adds an integer to the stream
func (stream *Stream) Push(value int) {
	if (*stream).len < 1 {
		return
	}
	if (*stream).k < (*stream).len {
		(*stream).MinHeap[(*stream).k] = value
		ind := (*stream).k
		for ind > 0 && (*stream).MinHeap[ind] < (*stream).MinHeap[(ind-1)/2] {
			tmp := (*stream).MinHeap[ind]
			(*stream).MinHeap[ind] = (*stream).MinHeap[(ind-1)/2]
			(*stream).MinHeap[(ind-1)/2] = tmp
			ind = (ind - 1) / 2
		}
		(*stream).k++
	} else if value > (*stream).MinHeap[0] {
		(*stream).MinHeap[0] = value
		ind := 0
		for {
			if 2*ind+2 < (*stream).len {
				if (*stream).MinHeap[2*ind+1] <= value &&
					(*stream).MinHeap[2*ind+1] <= (*stream).MinHeap[2*ind+2] {
					tmp := (*stream).MinHeap[2*ind+1]
					(*stream).MinHeap[2*ind+1] = value
					(*stream).MinHeap[ind] = tmp
					ind = 2*ind + 1
				} else if (*stream).MinHeap[2*ind+2] <= value &&
					(*stream).MinHeap[2*ind+2] <= (*stream).MinHeap[2*ind+1] {
					tmp := (*stream).MinHeap[2*ind+2]
					(*stream).MinHeap[2*ind+2] = value
					(*stream).MinHeap[ind] = tmp
					ind = 2*ind + 2
				} else {
					break
				}
			} else if 2*ind+1 < (*stream).len {
				if (*stream).MinHeap[2*ind+1] <= value {
					tmp := (*stream).MinHeap[2*ind+1]
					(*stream).MinHeap[2*ind+1] = value
					(*stream).MinHeap[ind] = tmp
					ind = 2*ind + 1
				} else {
					break
				}
			} else {
				break
			}
		}
	}
}

// Largest get at most `k` largest elements from the stream in arbitraty order
func (stream *Stream) Largest() []int {
	tmp := make([]int, (*stream).k)
	copy(tmp, (*stream).MinHeap)
	return tmp
}
