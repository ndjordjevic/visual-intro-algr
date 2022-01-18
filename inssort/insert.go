package inssort

func remEl(array []int, val int) []int {
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			array = append(array[:i], array[i+1:]...)
		}
	}

	return array
}

func insInArr(array []int, index int, value int) []int {
	array = append(array, 0)             // Making space for the new element
	copy(array[index+1:], array[index:]) // Shifting elements
	array[index] = value

	return array
}

func insertInSubArr(array []int, right int, value int) []int {
	array = remEl(array, value)

	for i := 0; i < right; i++ {
		if value < array[i] {
			array = insInArr(array, i, value) // Copying/inserting the value
			break
		}
	}

	return array
}

func insertSort(array []int) []int {
	var current int

	for i := 1; i < len(array); i++ {
		current = array[i]
		j := i - 1

		for j >= 0 && array[j] > current {
			array[j+1] = array[j]
			j = j - 1
		}

		array[j+1] = current
	}

	return array
}

func insertSort2(array []int) []int {
	for i := 1; i < len(array); i++ {
		for j := 0; j < i; j++ {
			if array[j] > array[i] {
				array[j], array[i] = array[i], array[j]
			}
		}
	}

	return array
}
