package ring_array_test

import (
	"snek/pkg/ring_array"
	"testing"
)

func expect(t *testing.T, a int, b int) {
	if a != b {
		t.Errorf("expected %d, got %d", a, b)
	}
}

func TestWrapping(t *testing.T) {
	r := ring_array.NewRingArray[int](5)

	r.PushFront(5)
	r.PushFront(2)
	r.PushFront(8)
	r.PushFront(4)

	var got int
	got, _ = r.PopBack()
	expect(t, 5, got)
	got, _ = r.PopBack()
	expect(t, 2, got)
	got, _ = r.PopBack()
	expect(t, 8, got)

	r.PushFront(3)
	r.PushFront(7)
	r.PushFront(12)

	got, _ = r.PopBack()
	expect(t, 4, got)
	got, _ = r.PopBack()
	expect(t, 3, got)
	got, _ = r.PopBack()
	expect(t, 7, got)
	got, _ = r.PopBack()
	expect(t, 12, got)
	got, _ = r.PopBack()
	expect(t, 0, got)
	r.PushFront(43)
	got, _ = r.PopBack()
	expect(t, 43, got)
}

func TestDirections(t *testing.T) {
	r := ring_array.NewRingArray[int](5)
	r.PushBack(4)
	r.PushBack(6)
	r.PushBack(2)

	var got int
	got, _ = r.PopFront()
	expect(t, 4, got)
	got, _ = r.PopFront()
	expect(t, 6, got)
	got, _ = r.PopFront()
	expect(t, 2, got)
}
