package streamlatest

// Stream accepts integers and allows to get any of `k` latest of them in O(1)
type Stream struct {
	start int
	len   int
	k     int
	arr   []int
}

// New creates a Stream instance
func New(k int) *Stream {
	if k < 1 {
		panic("invalid argument, expected to save at least one latest element")
	}
	return &(Stream{0, 0, k, make([]int, k)})
}

// Push adds an integer to the stream
func (s *Stream) Push(value int) {
	s.arr[s.start] = value
	s.start++
	if s.start == s.k {
		s.start = 0
	}
	if s.len != s.k {
		s.len++
	}
}

// Latest returns the `i`th latest element (zero-based) and a boolean that indicates success
func (s *Stream) Latest(i int) (int, bool) {
	if i < s.len && i > -1 {
		if s.start-1-i < 0 {
			return s.arr[s.k+(s.start-1-i)], true
		} else {
			return s.arr[s.start-1-i], true
		}
	} else {
		return -1, false
	}
}
