package matrix

import (
	"fmt"
	"math/rand"
	"time"
)

type UniqueMatrix struct {
	Rows      int
	Cols      int
	Data      [][]int
	UniqueMap map[int]bool
}

func NewUniqueMatrix(rows, cols int) *UniqueMatrix {
	return &UniqueMatrix{
		Rows:      rows,
		Cols:      cols,
		Data:      make([][]int, rows),
		UniqueMap: make(map[int]bool),
	}
}

func (m *UniqueMatrix) Generate(minVal, maxVal int) error {
	rand.Seed(time.Now().UnixNano())

	totalCells := m.Rows * m.Cols
	availableNumbers := maxVal - minVal + 1

	if totalCells > availableNumbers {
		return fmt.Errorf("недостаточно уникальных чисел: нужно %d, доступно %d",
			totalCells, availableNumbers)
	}

	for i := 0; i < m.Rows; i++ {
		m.Data[i] = make([]int, m.Cols)
		for j := 0; j < m.Cols; j++ {
			for {
				num := rand.Intn(availableNumbers) + minVal
				if !m.UniqueMap[num] {
					m.Data[i][j] = num
					m.UniqueMap[num] = true
					break
				}
			}
		}
	}

	return nil
}

func (m *UniqueMatrix) Print() {
	fmt.Println("\nСгенерированная матрица:")
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			fmt.Printf("%5d ", m.Data[i][j])
		}
		fmt.Println()
	}

	fmt.Printf("\nРазмер: %dx%d = %d элементов\n", m.Rows, m.Cols, m.Rows*m.Cols)
	fmt.Printf("Уникальных чисел: %d\n", len(m.UniqueMap))
}
