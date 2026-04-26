package main

import "fmt"

const arrSize int = 100

func sortBubble(arr *[arrSize]int) {
	for iteration := range arrSize - 1 {
		swapped := false
		for bubbleIndex := 0; bubbleIndex < arrSize-iteration-1; bubbleIndex++ {
			if arr[bubbleIndex] > arr[bubbleIndex+1] {
				arr[bubbleIndex], arr[bubbleIndex+1] = arr[bubbleIndex+1], arr[bubbleIndex]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func sortInsert(arr *[arrSize]int) {
	for indexInsertElement := 1; indexInsertElement < arrSize; indexInsertElement++ {
		insertElement := arr[indexInsertElement]
		sorted := indexInsertElement - 1
		for sorted >= 0 && insertElement < arr[sorted] {
			arr[sorted+1] = arr[sorted]
			sorted--
		}
		arr[sorted+1] = insertElement
	}
}

func sortMerge(arr *[arrSize]int) {
	temp := make([]int, arrSize)

	var mergeSort func(left, right int)
	mergeSort = func(left, right int) {
		if left >= right {
			return
		}

		mid := (left + right) / 2

		mergeSort(left, mid)
		mergeSort(mid+1, right)

		// Слияние
		i, j, k := left, mid+1, left

		for i <= mid && j <= right {
			if arr[i] <= arr[j] {
				temp[k] = arr[i]
				i++
			} else {
				temp[k] = arr[j]
				j++
			}
			k++
		}

		for i <= mid {
			temp[k] = arr[i]
			i++
			k++
		}

		for j <= right {
			temp[k] = arr[j]
			j++
			k++
		}

		for k := left; k <= right; k++ {
			arr[k] = temp[k]
		}
	}

	mergeSort(0, arrSize-1)
}

func task1() {
	arr := [arrSize]int{542, -565, 531, -294, -56, 14, 270, -51, -914, 605, -117, -768, 331, 708, -603, 84, -548, 579, 434, 751, 592, -349, 408, -602, 721, 909, 170, -432, -970, -171, -972, 316, 405, -676, -929, -795, -682, -646, 46, -609, -84, 180, -158, -662, -384, 854, -721, 39, 180, -197, -818, -946, -529, -555, -36, -853, -322, 540, -936, -919, 473, 978, 782, 586, 869, 333, -977, -548, -789, 988, -393, 807, -609, 997, 824, -480, -205, -576, 856, 494, 131, 40, -601, 467, 221, -640, 34, -220, 482, 948, 523, -27, -771, -914, 438, 957, 205, -411, -749, -723}
	p := &arr
	var commandsTask1 = `
	1 - Сортировка слиянием
	2 - Сортировка вставками
	3 - Сортирвка пузырьком
	4 - Исходный массив
	5 - выход в главное меню
	`
	var choiceTask1 int
loop:
	for {
		fmt.Print(commandsTask1)
		fmt.Scanf("%d", &choiceTask1)
		fmt.Scanln()

		switch choiceTask1 {
		case 1:
			fmt.Println("результат сортировки слиянием")
			sortMerge(p)
			fmt.Println(*p)
			break loop
		case 2:
			fmt.Println("результат сортировки вставками")
			sortInsert(p)
			fmt.Println(*p)
			break loop
		case 3:
			fmt.Println("результат сортировки пузыриком")
			sortBubble(p)
			fmt.Println(*p)
			break loop
		case 4:
			fmt.Println("исходный массив")
			fmt.Println(*p)
		case 5:
			break loop
		}
	}
}

func task2() {
	type Employee struct {
		Name     string // имя
		Age      int    // возраст
		Position string // позиция
		Salary   int    // зарплата
	}

	var commands = `
	1 - Добавить нового сотрудника
	2 - Удалить сотрудника
	3 - Вывести список сотрудников
	4 - выход в главное меню
	`
	const size = 512
	empls := [size]*Employee{}
	activeEmempls := 0
loop:
	for {
		cmd := 0
		fmt.Print(commands)
		fmt.Scanf("%d", &cmd)
		fmt.Scanln()

		switch cmd {
		case 1:
			if activeEmempls <= 512 {
				empl := new(Employee)
				fmt.Println("\nИмя:")
				fmt.Scanf("%s", &empl.Name)
				fmt.Scanln()
				fmt.Println("Возраст:")
				fmt.Scanf("%d", &empl.Age)
				fmt.Scanln()
				fmt.Println("Позиция:")
				fmt.Scanf("%s", &empl.Position)
				fmt.Scanln()
				fmt.Println("Зарплата:")
				fmt.Scanf("%d", &empl.Salary)
				fmt.Scanln()
				empls[activeEmempls] = empl
				activeEmempls++
			} else {
				fmt.Println("Список полон")
			}
		case 2:
			indexDelEmpl := 0
			if activeEmempls != 0 {
				fmt.Printf("Удаляем сотрудника (всего %d) № ", activeEmempls)
				fmt.Scanf("%d", &indexDelEmpl)
				fmt.Scanln()
				indexDelEmpl--

				for i := indexDelEmpl; i < activeEmempls-1; i++ {
					empls[i] = empls[i+1]
				}
				empls[activeEmempls-1] = nil
				activeEmempls--
				fmt.Printf("Сотрудник № %d удален, осталось %d ", indexDelEmpl+1, activeEmempls)
			} else {
				fmt.Printf("Список пуст")
			}
		case 3:
			fmt.Println("Вывод сотрудников")
			for indexEmpl := range activeEmempls {
				fmt.Printf("\nсотрудник № %d\nИмя: %s\nВозраст: %d\nПозиция: %s\nЗарплата: %d\n",
					indexEmpl+1, empls[indexEmpl].Name, empls[indexEmpl].Age,
					empls[indexEmpl].Position, empls[indexEmpl].Salary)

			}
		case 4:
			break loop
		}
	}
}

func main() {
	var commands = `
	1 - задание 1
	2 - задание 2
	3 - выход
	`
	var choice int
loop:
	for {
		fmt.Print(commands)
		fmt.Scanf("%d", &choice)
		fmt.Scanln()
		switch choice {
		case 1:
			task1()
		case 2:
			task2()
		case 3:
			break loop
		}
	}
}
