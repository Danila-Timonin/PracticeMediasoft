package main

type Stack struct {
	s    []any // слайс в котором хранятся значения в стеке
	head int   // индекс головы стека
}

func NewStack(size int) *Stack {
	return &Stack{
		s:    make([]any, size),
		head: -1,
	}
}

// Push - добавление в стек значения
func (s *Stack) Push(v any) {
	if s.head+1 >= len(s.s) {
		// Если стек заполнен, увеличиваем его размер
		newSlice := make([]any, len(s.s)*2)
		copy(newSlice, s.s)
		s.s = newSlice
	}
	s.head++
	s.s[s.head] = v
}

// Pop - получения значения из стека и его удаление из вершины
func (s *Stack) Pop() any {
	if s.head < 0 {
		return nil // стек пуст
	}
	v := s.s[s.head]
	s.s[s.head] = nil // освобождаем ссылку для GC
	s.head--
	return v
}

// Peek - просмотр значения на вершине стека
func (s *Stack) Peek() any {
	if s.head < 0 {
		return nil // стек пуст
	}
	return s.s[s.head]
}
