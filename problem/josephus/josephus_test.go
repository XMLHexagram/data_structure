package josephus

import (
	"testing"
)

func TestDealJosephus(t *testing.T) {
	var testStartAt int = 20
	var testInput []int = []int{3, 1, 7, 2, 4, 8, 4}

	testInputSlice := make([]int, 0, 100)
	for _, v := range testInput {
		testInputSlice = append(testInputSlice, v)
	}

	ll := CreateJosephus(testInputSlice)
	DealJosephus(testStartAt, ll)
}
