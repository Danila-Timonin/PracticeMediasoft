package linkedlist

import (
	"reflect"
	"testing"
)

func TestSinglyLinkedList(t *testing.T) {
	t.Run("Add и Values", func(t *testing.T) {
		l := NewSinglyLinkedList[int]()

		l.Add(10)
		l.Add(20)
		l.Add(30)

		expected := []int{10, 20, 30}
		result := l.Values()

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Values() = %v, expected %v", result, expected)
		}

		if l.Size != 3 {
			t.Errorf("Size = %d, expected 3", l.Size)
		}
	})

	t.Run("Get", func(t *testing.T) {
		l := NewSinglyLinkedList[string]()
		l.Add("a")
		l.Add("b")
		l.Add("c")

		tests := []struct {
			idx      int
			expected string
		}{
			{0, "a"},
			{1, "b"},
			{2, "c"},
		}

		for _, tt := range tests {
			if val := l.Get(tt.idx); val != tt.expected {
				t.Errorf("Get(%d) = %s, expected %s", tt.idx, val, tt.expected)
			}
		}
	})

	t.Run("Get with invalid index", func(t *testing.T) {
		l := NewSinglyLinkedList[int]()
		l.Add(10)
		l.Add(20)

		// Отрицательный индекс
		if val := l.Get(-1); val != 0 {
			t.Errorf("Get(-1) = %d, expected 0", val)
		}

		// Индекс больше размера
		if val := l.Get(5); val != 0 {
			t.Errorf("Get(5) = %d, expected 0", val)
		}
	})

	t.Run("Remove", func(t *testing.T) {
		l := NewSinglyLinkedList[int]()
		l.Add(10)
		l.Add(20)
		l.Add(30)
		l.Add(40)

		// Удаляем из середины
		l.Remove(1)
		expected := []int{10, 30, 40}
		if !reflect.DeepEqual(l.Values(), expected) {
			t.Errorf("After removing index 1: %v, expected %v", l.Values(), expected)
		}
		if l.Size != 3 {
			t.Errorf("Size = %d, expected 3", l.Size)
		}
	})

	t.Run("Remove first element", func(t *testing.T) {
		l := NewSinglyLinkedList[int]()
		l.Add(10)
		l.Add(20)
		l.Add(30)

		l.Remove(0)
		expected := []int{20, 30}
		if !reflect.DeepEqual(l.Values(), expected) {
			t.Errorf("After removing first: %v, expected %v", l.Values(), expected)
		}
	})

	t.Run("Remove last element", func(t *testing.T) {
		l := NewSinglyLinkedList[int]()
		l.Add(10)
		l.Add(20)
		l.Add(30)

		l.Remove(2)
		expected := []int{10, 20}
		if !reflect.DeepEqual(l.Values(), expected) {
			t.Errorf("After removing last: %v, expected %v", l.Values(), expected)
		}
	})

	t.Run("Remove single element", func(t *testing.T) {
		l := NewSinglyLinkedList[int]()
		l.Add(42)

		l.Remove(0)
		if l.Size != 0 {
			t.Errorf("Size = %d, expected 0", l.Size)
		}
		if l.First != nil {
			t.Error("First should be nil")
		}
		if l.Last != nil {
			t.Error("Last should be nil")
		}
	})

	t.Run("Remove with invalid index", func(t *testing.T) {
		l := NewSinglyLinkedList[int]()
		l.Add(10)
		l.Add(20)

		// Удаление с неверным индексом не должно менять список
		l.Remove(-1)
		if l.Size != 2 {
			t.Errorf("Size after invalid remove = %d, expected 2", l.Size)
		}

		l.Remove(10)
		if l.Size != 2 {
			t.Errorf("Size after invalid remove = %d, expected 2", l.Size)
		}
	})

	t.Run("Empty list", func(t *testing.T) {
		l := NewSinglyLinkedList[int]()

		if l.Size != 0 {
			t.Errorf("Size = %d, expected 0", l.Size)
		}

		if len(l.Values()) != 0 {
			t.Errorf("Values() = %v, expected empty slice", l.Values())
		}

		if val := l.Get(0); val != 0 {
			t.Errorf("Get(0) on empty list = %d, expected 0", val)
		}
	})

	t.Run("List with different types", func(t *testing.T) {
		// Список с float64
		l1 := NewSinglyLinkedList[float64]()
		l1.Add(1.1)
		l1.Add(2.2)
		if l1.Get(0) != 1.1 {
			t.Error("Float list failed")
		}

		// Список с bool
		l2 := NewSinglyLinkedList[bool]()
		l2.Add(true)
		l2.Add(false)
		if l2.Get(1) != false {
			t.Error("Bool list failed")
		}
	})
}
