package main

import (
	"fmt"
	"math/rand"
	"time"
)

func exapmle1() {
	fmt.Println("Стек")
	s := NewStack(3)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(*s)
	fmt.Println("Peek:", s.Peek())
	fmt.Println("Pop:", s.Pop())
	fmt.Println(*s)
	fmt.Println("Pop:", s.Pop())
	fmt.Println(*s)
	s.Push(4)
	fmt.Println("Push")
	fmt.Println(*s)
	fmt.Println("Peek:", s.Peek())

	fmt.Println("\nОчередь")
	q := NewQueue(3)
	q.Push("a")
	q.Push("b")
	q.Push("c")
	fmt.Println(*q)
	fmt.Println("Pop:", q.Pop())
	fmt.Println(*q)
	fmt.Println("Pop:", q.Pop())
	fmt.Println(*q)
	q.Push("d")
	fmt.Println("Pop:", q.Pop())
	fmt.Println("Pop:", q.Pop())
	fmt.Println(*q)

	fmt.Println("\nCписок")
	list := NewSinglyLinkedList()
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

func exapmle2() {
	fmt.Println("Конвертер римских чисел")
	converter := NewRomanConverter()

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

func exapmle3() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Генерация матрицы (максимум 200 элементов)")
	var rows, cols int
	var err bool

	for {
		fmt.Print("Число строк: ")
		fmt.Scanln()
		rows, err = getInt()
		if err {
			fmt.Println("Ошибка: введите положительное число")
			continue
		}
		fmt.Scanln()
		fmt.Print("Число столбцов: ")
		cols, err = getInt()
		if err {
			fmt.Println("Ошибка: введите положительное число")
			continue
		}
		if rows*cols > 200 {
			fmt.Printf("Ошибка: слишком много элементов! (%d > 200)\n", rows*cols)
			continue
		}

		break
	}

	matrix := make([][]int, rows)
	uniqueMap := make(map[int]bool)

	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			for {
				num := rand.Intn(601) - 300
				if !uniqueMap[num] {
					matrix[i][j] = num
					uniqueMap[num] = true
					break
				}
			}
		}
	}

	fmt.Println("\nСгенерированная матрица:")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%5d ", matrix[i][j])
		}
		fmt.Println()
	}

	fmt.Printf("\nРазмер: %dx%d = %d элементов\n", rows, cols, rows*cols)
	fmt.Printf("Уникальных чисел: %d\n", len(uniqueMap))
	fmt.Scanln()
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
			exapmle1()
		case 2:
			exapmle2()
		case 3:
			exapmle3()
		case 4:
			break loop
		}
	}
}
