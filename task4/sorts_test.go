package main

import (
	"reflect"
	"testing"
)

// Общий тест для всех сортировок
func TestSorts(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Пустой массив",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Массив из одного элемента",
			input:    []int{5},
			expected: []int{5},
		},
		{
			name:     "Уже отсортированный массив",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Обратно отсортированный массив",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Массив с дубликатами",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3},
			expected: []int{1, 1, 2, 3, 3, 4, 5, 5, 6, 9},
		},
		{
			name:     "Массив с отрицательными числами",
			input:    []int{-5, 3, -2, 0, -1, 4, -3},
			expected: []int{-5, -3, -2, -1, 0, 3, 4},
		},
		{
			name:     "Массив из одинаковых чисел",
			input:    []int{7, 7, 7, 7, 7},
			expected: []int{7, 7, 7, 7, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Тестируем сортировку пузырьком
			t.Run("BubbleSort", func(t *testing.T) {
				arr := make([]int, len(tt.input))
				copy(arr, tt.input)
				sortBubble(arr)
				if !reflect.DeepEqual(arr, tt.expected) {
					t.Errorf("BubbleSort() = %v, expected %v", arr, tt.expected)
				}
			})

			// Тестируем сортировку вставками
			t.Run("InsertSort", func(t *testing.T) {
				arr := make([]int, len(tt.input))
				copy(arr, tt.input)
				sortInsert(arr)
				if !reflect.DeepEqual(arr, tt.expected) {
					t.Errorf("InsertSort() = %v, expected %v", arr, tt.expected)
				}
			})

			// Тестируем сортировку слиянием
			t.Run("MergeSort", func(t *testing.T) {
				arr := make([]int, len(tt.input))
				copy(arr, tt.input)
				sortMerge(arr)
				if !reflect.DeepEqual(arr, tt.expected) {
					t.Errorf("MergeSort() = %v, expected %v", arr, tt.expected)
				}
			})
		})
	}
}
