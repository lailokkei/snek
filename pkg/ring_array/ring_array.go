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

func (r *RingArray[T]) Head() (T, error) {
	if r.Length <= 0 {
		var head T
		return head, errors.New("no elements")
	}
	return r.array[r.head], nil
}

func (r *RingArray[T]) Tail() (T, error) {
	if r.Length <= 0 {
		var tail T
		return tail, errors.New("no elements")
	}
	return r.array[r.tail], nil
}

func (r *RingArray[T]) Array() []T {
	array := make([]T, r.Length)
	idx := 0
	if r.head < r.tail {
		for i := r.tail; i < r.size; i++ {
			array[idx] = r.array[i]
			idx++
		}
		for i := 0; i < r.head+1; i++ {
			array[idx] = r.array[i]
			idx++
		}
		return array
	}

	for i := r.tail; i < r.head+1; i++ {
		array[idx] = r.array[i]
		idx++
	}

	return array
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
