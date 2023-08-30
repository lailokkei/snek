package ringbuffer

import (
	"errors"
)

type RingBuffer[T any] struct {
	buffer []T
	head   int
	tail   int
	size   int
	Length int
}

func NewRingBuffer[T any](size int) RingBuffer[T] {
	return RingBuffer[T]{
		buffer: make([]T, size),
		head:   0,
		tail:   0,
		size:   size,
		Length: 0,
	}
}

func (r RingBuffer[T]) wrapIndex(i int) int {
	return i % r.size
}

func (r *RingBuffer[T]) Push(value T) error {
	if r.Length >= len(r.buffer) {
		return errors.New("Buffer is full.")
	}

	r.Length++
	r.buffer[r.tail] = value
	r.tail = r.wrapIndex(r.tail + 1)

	return nil
}

func (r *RingBuffer[T]) Pop() (T, error) {
	if r.Length <= 0 {
		var t T
		return t, errors.New("Buffer is empty")
	}

	r.Length--
	value := r.buffer[r.head]
	r.head = r.wrapIndex(r.head + 1)

	return value, nil
}
