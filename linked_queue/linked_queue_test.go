package linked_queue

import (
	"testing"
)

func TestQueue_IsEmpty(t *testing.T) {
	q := NewQueue()

	if q.IsEmpty() != true {
		t.Error("empty to non-empty")
	}
}

func TestQueue_Push(t *testing.T) {
	q := NewQueue()
	q.Push(1)
	if q.Bottom.Data != 1 {
		t.Error("wrong bottom node")
	}
	q.Push(2)
	if q.Bottom.Data != 2 {
		t.Error("wrong bottom node")
	}
	q.Push(3)
	if q.Bottom.Data != 3 {
		t.Error("wrong bottom node")
	}
	if q.Size() != 3 {
		t.Error("wrong list length")
	}
}

func TestQueue_Pop(t *testing.T) {
	q := NewQueue()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	if q.Pop().Data != 1 {
		t.Error("queue error: wrong node poped")
	}
	if q.Pop().Data != 2 {
		t.Error("queue error: wrong node poped")
	}
	if q.Pop().Data != 3 {
		t.Error("queue error: wrong node poped")
	}
	if q.Size() != 0 {
		t.Error("queue error: wrong size")
	}
}