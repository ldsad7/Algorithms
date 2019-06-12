package queue

import (
	"fmt"
)

type Object struct {
	Value uint32
	Max   uint32
}

// MaxQueue is a FIFO queue that allows fast queries for the maximum among currently enqueued elements.
type MaxQueue struct {
	Input  []Object
	Output []Object
}

// New creates an instance of MaxQueue
func New() *MaxQueue {
	return new(MaxQueue)
}

// Push inserts an element to the tail
func (q *MaxQueue) Push(value uint32) {
	max := value
	if len(q.Input) > 0 {
		tmp := q.Input[len(q.Input)-1].Max
		if tmp > max {
			max = tmp
		}
	}
	q.Input = append(q.Input, Object{value, max})
}

// Pop removes an element from the head
func (q *MaxQueue) Pop() (uint32, error) {
	if len(q.Output) > 0 {
		tmp := q.Output[len(q.Output)-1]
		q.Output = q.Output[:len(q.Output)-1]
		return tmp.Value, nil
	} else if len(q.Input) > 0 {
		for j := len(q.Input) - 1; j >= 0; j-- {
			max := q.Input[j].Value
			if len(q.Output) > 0 && q.Output[len(q.Output)-1].Max > max {
				max = q.Output[len(q.Output)-1].Max
			}
			q.Output = append(q.Output, Object{q.Input[j].Value, max})
		}
		q.Input = make([]Object, 0)
		return q.Pop()
	} else {
		return 0, fmt.Errorf("Pop Error")
	}
}

// Max returns maximum among currently enqueued elements
func (q *MaxQueue) Max() (uint32, error) {
	if len(q.Input) == 0 && len(q.Output) == 0 {
		return 0, fmt.Errorf("Max Error")
	} else if len(q.Input) > 0 && len(q.Output) > 0 {
		if q.Input[len(q.Input)-1].Max > q.Output[len(q.Output)-1].Max {
			return q.Input[len(q.Input)-1].Max, nil
		}
		return q.Output[len(q.Output)-1].Max, nil
	} else if len(q.Input) > 0 {
		return q.Input[len(q.Input)-1].Max, nil
	} else {
		return q.Output[len(q.Output)-1].Max, nil
	}
}
