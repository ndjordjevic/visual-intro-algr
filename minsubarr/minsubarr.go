package minsubarr

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