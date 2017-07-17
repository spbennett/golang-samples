package sample

import "testing"
import "reflect"

var listOne = []int{3, 4, 6, 10, 11, 15}
var listTwo = []int{1, 5, 8, 12, 14, 19}
var expectedMergeOneTwo = []int{1, 3, 4, 5, 6, 8, 10, 11, 12, 14, 15, 19}

func TestMergeLists(t *testing.T) {
	actualList := MergeLists(listOne, listTwo)

	if !reflect.DeepEqual(actualList, expectedMergeOneTwo) {
		t.Error(
			"For", listOne,
			"and", listTwo,
			"expected", expectedMergeOneTwo,
			"got", actualList,
		)
	}

}
