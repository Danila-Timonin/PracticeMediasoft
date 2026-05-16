package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("Push и Pop", func(t *testing.T) {
		s := NewStack[int](3)

		// Проверка Push
		s.Push(1)
		s.Push(2)
		s.Push(3)

		if s.Size() != 3 {
			t.Errorf("Size() = %d, expected 3", s.Size())
		}

		// Проверка Pop
		if val := s.Pop(); val != 3 {
			t.Errorf("Pop() = %d, expected 3", val)
		}
		if val := s.Pop(); val != 2 {
			t.Errorf("Pop() = %d, expected 2", val)
		}
		if val := s.Pop(); val != 1 {
			t.Errorf("Pop() = %d, expected 1", val)
		}

		if !s.IsEmpty() {
			t.Error("Stack should be empty")
		}
	})

	t.Run("Peek", func(t *testing.T) {
		s := NewStack[string](3)
		s.Push("hello")
		s.Push("world")

		if val := s.Peek(); val != "world" {
			t.Errorf("Peek() = %s, expected world", val)
		}
		if s.Size() != 2 {
			t.Errorf("Size() after Peek = %d, expected 2", s.Size())
		}
	})

	t.Run("Pop on empty stack", func(t *testing.T) {
		s := NewStack[int](3)
		val := s.Pop()
		if val != 0 {
			t.Errorf("Pop() on empty stack = %d, expected 0", val)
		}
	})

	t.Run("Automatic resize", func(t *testing.T) {
		s := NewStack[int](2)
		s.Push(1)
		s.Push(2)
		s.Push(3) // Should resize

		if s.Size() != 3 {
			t.Errorf("Size() after resize = %d, expected 3", s.Size())
		}

		if val := s.Pop(); val != 3 {
			t.Errorf("Pop() = %d, expected 3", val)
		}
	})

	t.Run("Stack with different types", func(t *testing.T) {
		// Стек со строками
		s1 := NewStack[string](2)
		s1.Push("a")
		s1.Push("b")
		if s1.Pop() != "b" {
			t.Error("String stack failed")
		}

		// Стек с float64
		s2 := NewStack[float64](2)
		s2.Push(1.1)
		s2.Push(2.2)
		if s2.Pop() != 2.2 {
			t.Error("Float stack failed")
		}
	})
}
