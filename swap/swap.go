package swap

func swap(list []int, start int, end int) {
	temp := list[start]
	list[start] = list[end]
	list[end] = temp
}