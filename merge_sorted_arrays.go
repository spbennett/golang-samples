package sample

// MergeLists takes two sorted lists and returns one combined sorted list.
func MergeLists(listOne []int, listTwo []int) []int {
	indexOne := 0
	indexTwo := 0

	capacity := len(listOne) + len(listTwo)
	combinedList := make([]int, capacity)

	for i := 0; i < len(combinedList); i++ {
		if indexTwo == len(listTwo) || indexOne < len(listOne) && listOne[indexOne] < listTwo[indexTwo] {
			combinedList[i] = listOne[indexOne]
			indexOne++
		} else {
			combinedList[i] = listTwo[indexTwo]
			indexTwo++
		}
	}

	return combinedList
}
