package main

// Сортировка пузырьком
func sortBubble(arr []int) {
	size := len(arr)
	for iteration := 0; iteration < size-1; iteration++ {
		swapped := false
		for bubbleIndex := 0; bubbleIndex < size-iteration-1; bubbleIndex++ {
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

// Сортировка вставками
func sortInsert(arr []int) {
	size := len(arr)
	for indexInsertElement := 1; indexInsertElement < size; indexInsertElement++ {
		insertElement := arr[indexInsertElement]
		sorted := indexInsertElement - 1
		for sorted >= 0 && insertElement < arr[sorted] {
			arr[sorted+1] = arr[sorted]
			sorted--
		}
		arr[sorted+1] = insertElement
	}
}

// Сортировка слиянием
func sortMerge(arr []int) {
	size := len(arr)
	temp := make([]int, size)

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

	mergeSort(0, size-1)
}
