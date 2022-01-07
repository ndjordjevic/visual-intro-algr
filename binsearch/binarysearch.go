package binsearch

func binarySearch(list []int, target int) int {
	min := 0
	max := len(list) - 1
	var guess int

	for min <= max {
		guess = (min + max) / 2

		if list[guess] == target {
			return guess
		}

		if list[guess] < target {
			min = guess + 1
		} else {
			max = guess - 1
		}
	}

	return -1
}
