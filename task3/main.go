package main

import (
	"fmt"
	"math/rand"
	"task3/linkedlist"
	"task3/matrix"
	"task3/queue"
	"task3/roman"
	"task3/stack"
	"time"
)

func example1() {
	fmt.Println("Стек")
	s := stack.NewStack[int](3) // Явно указываем тип int
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s)
	fmt.Println("Peek:", s.Peek())
	fmt.Println("Pop:", s.Pop())
	fmt.Println(s)
	fmt.Println("Pop:", s.Pop())
	fmt.Println(s)
	s.Push(4)
	fmt.Println("Push")
	fmt.Println(s)
	fmt.Println("Peek:", s.Peek())

	fmt.Println("\nОчередь")
	q := queue.NewQueue[string](3) // Явно указываем тип string
	q.Push("a")
	q.Push("b")
	q.Push("c")
	fmt.Println(q)
	fmt.Println("Pop:", q.Pop())
	fmt.Println(q)
	fmt.Println("Pop:", q.Pop())
	fmt.Println(q)
	q.Push("d")
	fmt.Println("Pop:", q.Pop())
	fmt.Println("Pop:", q.Pop())
	fmt.Println(q)

	fmt.Println("\nCписок")
	list := linkedlist.NewSinglyLinkedList[int]() // Явно указываем тип int
	list.Add(10)
	list.Add(20)
	list.Add(30)
	fmt.Println(list.Values())
	fmt.Println("Get(1):", list.Get(1))
	fmt.Println("Values:", list.Values())
	fmt.Println("Remove(1)")
	list.Remove(1)
	fmt.Println("After remove:", list.Values())
}

func example2() {
	fmt.Println("Конвертер римских чисел")
	converter := roman.NewRomanConverter()

	testNumbers := []string{
		"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI",
		"XII", "XIII", "XIV", "XIX", "XX", "XL", "L", "XC", "C", "CD", "D", "CM", "M",
		"MCMXCIV", "MMXXIV", "MMMCMXCIX",
	}

	for _, roman := range testNumbers {
		if arabic, err := converter.RomanToArabic(roman); err == nil {
			fmt.Printf("%-10s -> %d\n", roman, arabic)
		} else {
			fmt.Printf("%-10s -> Ошибка: %v\n", roman, err)
		}
	}
}

func getInt() (int, bool) {
	var num int
	fmt.Scanf("%d", &num)
	return num, num <= 0
}

func example3() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Генерация матрицы (максимум 200 элементов)")
	var rows, cols int

	for {
		fmt.Print("Число строк: ")
		fmt.Scanln(&rows) // Читаем напрямую в переменную
		if rows > 0 {
			break
		}
		fmt.Println("Ошибка: введите положительное число")
	}

	for {
		fmt.Print("Число столбцов: ")
		fmt.Scanln(&cols) // Читаем напрямую в переменную
		if cols > 0 {
			break
		}
		fmt.Println("Ошибка: введите положительное число")
	}

	// Проверка количества элементов
	for rows*cols > 200 {
		fmt.Printf("Ошибка: слишком много элементов! (%d > 200)\n", rows*cols)

		for {
			fmt.Print("Число строк: ")
			fmt.Scanln(&rows)
			if rows > 0 {
				break
			}
			fmt.Println("Ошибка: введите положительное число")
		}

		for {
			fmt.Print("Число столбцов: ")
			fmt.Scanln(&cols)
			if cols > 0 {
				break
			}
			fmt.Println("Ошибка: введите положительное число")
		}
	}

	m := matrix.NewUniqueMatrix(rows, cols)
	m.Generate(-300, 300)

	fmt.Println("\nСгенерированная матрица:")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%5d ", m.Data[i][j])
		}
		fmt.Println()
	}

	fmt.Printf("\nРазмер: %dx%d = %d элементов\n", rows, cols, rows*cols)
	fmt.Printf("Уникальных чисел: %d\n", len(m.UniqueMap))
}

func main() {
	var commands = `
	1 - Пример работы со структурами
	2 - Конвертер 
	3 - Двумерный массив
	4 - выход
	`
	var choice int
loop:
	for {
		fmt.Print(commands)
		fmt.Scanf("%d", &choice)
		fmt.Scanln()
		switch choice {
		case 1:
			example1()
		case 2:
			example2()
		case 3:
			example3()
		case 4:
			break loop
		}
	}
}
