package mergesort

func mergeSort(array []float64) {
	arrayLenght := len(array)
	if arrayLenght <= 1 {
		return
	}

	middleArray := arrayLenght / 2
	leftArray := append([]float64{}, array[:middleArray]...)
	rightArray := append([]float64{}, array[middleArray:]...)

	mergeSort(leftArray)
	mergeSort(rightArray)

	merge(array, leftArray, rightArray)
}

func merge(array, leftArray, rightArray []float64) {
	leftArrayLenght := len(leftArray)
	rightArrayLenght := len(rightArray)

	i, j, k := 0, 0, 0
	for i < leftArrayLenght && j < rightArrayLenght {
		if leftArray[i] < rightArray[j] {
			array[k] = leftArray[i]
			i++
		} else {
			array[k] = rightArray[j]
			j++
		}
		k++
	}

	for i < leftArrayLenght {
		array[k] = leftArray[i]
		i++
		k++
	}

	for j < rightArrayLenght {
		array[k] = rightArray[j]
		j++
		k++
	}
}
