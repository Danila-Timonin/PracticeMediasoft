package stack

type Stack[T any] struct {
	s    []T
	head int
}

func NewStack[T any](size int) *Stack[T] {
	return &Stack[T]{
		s:    make([]T, size),
		head: -1,
	}
}

func (s *Stack[T]) Push(v T) {
	if s.head+1 >= len(s.s) {
		newSlice := make([]T, len(s.s)*2)
		copy(newSlice, s.s)
		s.s = newSlice
	}
	s.head++
	s.s[s.head] = v
}

func (s *Stack[T]) Pop() T {
	if s.head < 0 {
		var zero T
		return zero
	}
	v := s.s[s.head]
	s.s[s.head] = *new(T) // освобождаем ссылку
	s.head--
	return v
}

func (s *Stack[T]) Peek() T {
	if s.head < 0 {
		var zero T
		return zero
	}
	return s.s[s.head]
}

func (s *Stack[T]) IsEmpty() bool {
	return s.head < 0
}

func (s *Stack[T]) Size() int {
	return s.head + 1
}
