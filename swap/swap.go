package swap

func swap(list []int, start int, end int) {
	list[start], list[end] = list[end], list[start]
}