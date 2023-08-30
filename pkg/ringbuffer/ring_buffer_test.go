package ringbuffer_test

import (
	"snek/pkg/ringbuffer"
	"testing"
)

func expect(t *testing.T, a int, b int) {
	if a != b {
		t.Errorf("expected %d, got %d", a, b)
	}
}

func TestRingBuffer(t *testing.T) {
	r := ringbuffer.NewRingBuffer[int](5)

	r.Push(5)
	r.Push(2)
	r.Push(8)
	r.Push(4)

	var got int
	got, _ = r.Pop()
	expect(t, 5, got)
	got, _ = r.Pop()
	expect(t, 2, got)
	got, _ = r.Pop()
	expect(t, 8, got)

	r.Push(3)
	r.Push(7)
	r.Push(12)

	got, _ = r.Pop()
	expect(t, 4, got)
	got, _ = r.Pop()
	expect(t, 3, got)
	got, _ = r.Pop()
	expect(t, 7, got)
	got, _ = r.Pop()
	expect(t, 12, got)
}
