package quicksort

func quickSort(array []int) {
	quickSortFunctiion(array, 0, len(array)-1)
}

func quickSortFunctiion(array []int, low int, high int) {
	if low >= high {
		return
	}

	var pivotIndex = partion(array, low, high)

	quickSortFunctiion(array, low, pivotIndex-1)
	quickSortFunctiion(array, pivotIndex+1, high)
}

func partion(array []int, low int, high int) int {
	pivot := array[high]

	index := low - 1

	for j := low; j <= high-1; j++ {
		if array[j] <= pivot {
			index++
			array[index], array[j] = array[j], array[index]
		}
	}
	index++
	array[high], array[index] = array[index], array[high]
	return index
}
