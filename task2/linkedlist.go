package main

type Item struct {
	V    any
	Next *Item
}

type SinglyLinkedList struct {
	First *Item
	Last  *Item
	Size  int
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

// Add - добавление значения в связный список
func (l *SinglyLinkedList) Add(v any) {
	newItem := &Item{V: v}

	if l.First == nil {
		l.First = newItem
		l.Last = newItem
	} else {
		l.Last.Next = newItem
		l.Last = newItem
	}
	l.Size++
}

// Get - получение значения по индексу из связанного списка
func (l *SinglyLinkedList) Get(idx int) any {
	if idx < 0 || idx >= l.Size {
		return nil // индекс вне диапазона
	}

	current := l.First
	for i := 0; i < idx; i++ {
		current = current.Next
	}
	return current.V
}

// Remove - удаление значения по индексу из списка
func (l *SinglyLinkedList) Remove(idx int) {
	if idx < 0 || idx >= l.Size {
		return // индекс вне диапазона
	}

	if idx == 0 {
		// Удаляем первый элемент
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
			// Удаляем последний элемент
			l.Last = current
		}
	}
	l.Size--
}

// Values - получение слайса значений из списка
func (l *SinglyLinkedList) Values() []any {
	result := make([]any, l.Size)
	current := l.First
	for i := 0; i < l.Size; i++ {
		result[i] = current.V
		current = current.Next
	}
	return result
}
