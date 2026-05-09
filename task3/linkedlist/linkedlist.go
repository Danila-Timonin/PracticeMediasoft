package linkedlist

type Item[T any] struct {
	V    T
	Next *Item[T]
}

type SinglyLinkedList[T any] struct {
	First *Item[T]
	Last  *Item[T]
	Size  int
}

func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

func (l *SinglyLinkedList[T]) Add(v T) {
	newItem := &Item[T]{V: v}

	if l.First == nil {
		l.First = newItem
		l.Last = newItem
	} else {
		l.Last.Next = newItem
		l.Last = newItem
	}
	l.Size++
}

func (l *SinglyLinkedList[T]) Get(idx int) T {
	var zero T
	if idx < 0 || idx >= l.Size {
		return zero
	}

	current := l.First
	for i := 0; i < idx; i++ {
		current = current.Next
	}
	return current.V
}

func (l *SinglyLinkedList[T]) Remove(idx int) {
	if idx < 0 || idx >= l.Size {
		return
	}

	if idx == 0 {
		l.First = l.First.Next
		if l.Size == 1 {
			l.Last = nil
		}
	} else {
		current := l.First
		for i := 0; i < idx-1; i++ {
			current = current.Next
		}

		current.Next = current.Next.Next
		if idx == l.Size-1 {
			l.Last = current
		}
	}
	l.Size--
}

func (l *SinglyLinkedList[T]) Values() []T {
	result := make([]T, l.Size)
	current := l.First
	for i := 0; i < l.Size; i++ {
		result[i] = current.V
		current = current.Next
	}
	return result
}
