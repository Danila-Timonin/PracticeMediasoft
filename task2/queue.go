package main

type Queue struct {
	s         []any // слайс в котором хранятся значения
	low, high int   // индексы верхней и нижней границы очереди
	size      int   // размер очереди
}

func NewQueue(size int) *Queue {
	return &Queue{
		s:    make([]any, size),
		size: size,
		low:  -1,
		high: -1,
	}
}

// Push - добавление в очередь значения
func (q *Queue) Push(v any) {
	if q.low == -1 && q.high == -1 {
		// Очередь пуста
		q.low = 0
		q.high = 0
		q.s[0] = v
		return
	}

	nextHigh := (q.high + 1) % q.size
	if nextHigh == q.low {
		// Очередь заполнена, увеличиваем размер
		newSize := q.size * 2
		newSlice := make([]any, newSize)

		// Копируем элементы в новый слайс
		for i := 0; i < q.size; i++ {
			idx := (q.low + i) % q.size
			newSlice[i] = q.s[idx]
		}

		q.s = newSlice
		q.low = 0
		q.high = q.size - 1
		q.size = newSize
		nextHigh = q.high + 1
	}

	q.high = nextHigh
	q.s[q.high] = v
}

// Pop - получения значения из очереди и его удаление
func (q *Queue) Pop() any {
	if q.low == -1 {
		return nil // очередь пуста
	}

	v := q.s[q.low]
	q.s[q.low] = nil // освобождаем ссылку для GC

	if q.low == q.high {
		// В очереди был только один элемент
		q.low = -1
		q.high = -1
	} else {
		q.low = (q.low + 1) % q.size
	}

	return v
}
