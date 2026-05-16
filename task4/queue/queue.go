package queue

type Queue[T any] struct {
	s         []T
	low, high int
	size      int
}

func NewQueue[T any](size int) *Queue[T] {
	return &Queue[T]{
		s:    make([]T, size),
		size: size,
		low:  -1,
		high: -1,
	}
}

func (q *Queue[T]) Push(v T) {
	if q.low == -1 && q.high == -1 {
		q.low = 0
		q.high = 0
		q.s[0] = v
		return
	}

	nextHigh := (q.high + 1) % q.size
	if nextHigh == q.low {
		newSize := q.size * 2
		newSlice := make([]T, newSize)

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

func (q *Queue[T]) Pop() T {
	if q.low == -1 {
		var zero T
		return zero
	}

	v := q.s[q.low]
	var zero T
	q.s[q.low] = zero

	if q.low == q.high {
		q.low = -1
		q.high = -1
	} else {
		q.low = (q.low + 1) % q.size
	}

	return v
}

func (q *Queue[T]) IsEmpty() bool {
	return q.low == -1
}

func (q *Queue[T]) Size() int {
	if q.IsEmpty() {
		return 0
	}
	if q.high >= q.low {
		return q.high - q.low + 1
	}
	return q.size - q.low + q.high + 1
}
