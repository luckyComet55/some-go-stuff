package circlebuffer

import (
	"errors"
	"sync"
)

// Потокобезопасный кольцевой буффер
type CircleBuffer[T any] struct {
	buffer   []T
	capacity int
	size     int
	start    int
	offset   int
	mtx      sync.Mutex
}

func NewCircleBuffer[T any](initCap int) *CircleBuffer[T] {
	return &CircleBuffer[T]{
		buffer:   make([]T, initCap),
		capacity: initCap,
		size:     0,
		start:    0,
		offset:   0,
		mtx:      sync.Mutex{},
	}
}

func (buffer *CircleBuffer[T]) GetSize() int {
	buffer.mtx.Lock()
	defer buffer.mtx.Unlock()

	return buffer.size
}

func (buffer *CircleBuffer[T]) AddElem(elem T) bool {
	buffer.mtx.Lock()
	defer buffer.mtx.Unlock()

	if (buffer.offset+1)%buffer.capacity == buffer.start {
		return false
	}

	if buffer.offset+1 == buffer.capacity {
		buffer.offset = 0
	} else {
		buffer.offset++
	}
	buffer.buffer[buffer.offset] = elem
	buffer.size++
	return true
}

func (buffer *CircleBuffer[T]) PopFront() (T, error) {
	buffer.mtx.Lock()
	defer buffer.mtx.Unlock()

	if buffer.size == 0 {
		var res T
		return res, errors.New("empty buffer")
	}

	idx := buffer.start
	if (buffer.start + 1) == buffer.capacity {
		buffer.start = 0
	}
	return buffer.buffer[idx], nil
}
