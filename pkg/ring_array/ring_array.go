package ring_array

import (
	"errors"
)

type RingArray[T any] struct {
	array  []T
	head   int
	tail   int
	size   int
	Length int
}

func NewRingArray[T any](size int) RingArray[T] {
	return RingArray[T]{
		array:  make([]T, size),
		head:   0,
		tail:   0,
		size:   size,
		Length: 0,
	}
}

func (r RingArray[T]) wrapIndex(i int) int {
	if i < 0 {
		return r.size - ((i * -1) % r.size)
	}
	return i % r.size
}

func (r *RingArray[T]) Head() T {
	return r.array[r.head]
}

func (r *RingArray[T]) Tail() T {
	return r.array[r.tail]
}

func (r *RingArray[T]) PushFront(value T) error {
	if r.Length >= r.size {
		return errors.New("Buffer is full.")
	}

	if r.Length > 0 {
		r.head = r.wrapIndex(r.head + 1)
	}

	r.Length++
	r.array[r.head] = value

	return nil
}

func (r *RingArray[T]) PushBack(value T) error {
	if r.Length >= r.size {
		return errors.New("Buffer is full.")
	}

	if r.Length > 0 {
		r.tail = r.wrapIndex(r.tail - 1)
	}

	r.Length++
	r.array[r.tail] = value

	return nil
}

func (r *RingArray[T]) PopBack() (T, error) {
	var value T
	if r.Length <= 0 {
		return value, errors.New("Buffer is empty")
	}

	r.Length--
	value = r.array[r.tail]

	if r.Length > 0 {
		r.tail = r.wrapIndex(r.tail + 1)
	}

	return value, nil
}

func (r *RingArray[T]) PopFront() (T, error) {
	var value T
	if r.Length <= 0 {
		return value, errors.New("Buffer is empty")
	}

	r.Length--
	value = r.array[r.head]

	if r.Length > 0 {
		r.head = r.wrapIndex(r.head - 1)
	}

	return value, nil
}
