package selsort

func swap(list []int, start int, end int) {
	temp := list[start]
	list[start] = list[end]
	list[end] = temp
}

func indexOfMinimum(list []int, startI int) int {
	minV := list[startI]
	minI := startI

	for i := startI; i < len(list); i++ {
		if list[i] < minV {
			minV = list[i]
			minI = i
		}
	}

	return minI
}

func selSort(list []int) []int {
	for i := 0; i < len(list); i++ {
		end := indexOfMinimum(list, i)
		if end > i {
			swap(list, i, end)
		}
	}

	return list
}
