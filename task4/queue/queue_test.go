package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	t.Run("Push и Pop", func(t *testing.T) {
		q := NewQueue[string](3)

		// Проверка Push
		q.Push("a")
		q.Push("b")
		q.Push("c")

		if q.Size() != 3 {
			t.Errorf("Size() = %d, expected 3", q.Size())
		}

		// Проверка Pop (FIFO)
		if val := q.Pop(); val != "a" {
			t.Errorf("Pop() = %s, expected a", val)
		}
		if val := q.Pop(); val != "b" {
			t.Errorf("Pop() = %s, expected b", val)
		}
		if val := q.Pop(); val != "c" {
			t.Errorf("Pop() = %s, expected c", val)
		}

		if !q.IsEmpty() {
			t.Error("Queue should be empty")
		}
	})

	t.Run("Push after Pop", func(t *testing.T) {
		q := NewQueue[int](3)
		q.Push(1)
		q.Push(2)
		q.Push(3)

		q.Pop() // удаляем 1
		q.Pop() // удаляем 2

		q.Push(4)
		q.Push(5)

		if val := q.Pop(); val != 3 {
			t.Errorf("Pop() = %d, expected 3", val)
		}
		if val := q.Pop(); val != 4 {
			t.Errorf("Pop() = %d, expected 4", val)
		}
		if val := q.Pop(); val != 5 {
			t.Errorf("Pop() = %d, expected 5", val)
		}
	})

	t.Run("Circular behavior", func(t *testing.T) {
		q := NewQueue[int](3)
		q.Push(1)
		q.Push(2)
		q.Push(3)

		q.Pop()
		q.Push(4)

		if val := q.Pop(); val != 2 {
			t.Errorf("Pop() = %d, expected 2", val)
		}
		if val := q.Pop(); val != 3 {
			t.Errorf("Pop() = %d, expected 3", val)
		}
		if val := q.Pop(); val != 4 {
			t.Errorf("Pop() = %d, expected 4", val)
		}
	})

	t.Run("Automatic resize", func(t *testing.T) {
		q := NewQueue[int](2)
		q.Push(1)
		q.Push(2)
		q.Push(3)

		if q.Size() != 3 {
			t.Errorf("Size() after resize = %d, expected 3", q.Size())
		}

		if val := q.Pop(); val != 1 {
			t.Errorf("Pop() = %d, expected 1", val)
		}
		if val := q.Pop(); val != 2 {
			t.Errorf("Pop() = %d, expected 2", val)
		}
		if val := q.Pop(); val != 3 {
			t.Errorf("Pop() = %d, expected 3", val)
		}
	})

	t.Run("Pop on empty queue", func(t *testing.T) {
		q := NewQueue[int](3)
		val := q.Pop()
		if val != 0 {
			t.Errorf("Pop() on empty queue = %d, expected 0", val)
		}
	})

	t.Run("Queue with different types", func(t *testing.T) {
		// Очередь с целыми числами
		q1 := NewQueue[int](2)
		q1.Push(10)
		q1.Push(20)
		if q1.Pop() != 10 {
			t.Error("Int queue failed")
		}

		// Очередь с булевыми значениями
		q2 := NewQueue[bool](2)
		q2.Push(true)
		q2.Push(false)
		if q2.Pop() != true {
			t.Error("Bool queue failed")
		}
	})
}
