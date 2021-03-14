package multiplicationTable

func MultiplicationTable(size int) [][]int {

	var resultSlice = make([][]int, size)

	for l1 := 0; l1 < size; l1++ {
		for l2 := 1; l2 <= size; l2++ {
			resultSlice[l1] = append(resultSlice[l1], l2*(l1+1))
		}
	}
	return resultSlice
}
