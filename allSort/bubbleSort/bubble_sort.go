package bubblesort

func bubbleSort(array []float64) {
	lenght := len(array)

	for i := 0; i < lenght; i++ {
		for j := 0; j < lenght-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

}
