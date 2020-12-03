package josephus

import (
	"github.com/lmx-Hexagram/data_structure/link/circularLinkedList"
	"testing"
)

func TestDealJosephus(t *testing.T) {
	var testStartAt int = 20
	var testInput []int = []int{3, 1, 7, 2, 4, 8, 4}

	ll := CreateJosephus(testInput)
	DealJosephus(testStartAt, ll)
}

func CreateJosephus(testInput []int) *circularLinkedList.CNode {
	ll := circularLinkedList.Init()
	for i := 1; i <= len(testInput); i++ {
		ll.AddAtEnd(Data{
			ID:   i,
			Code: testInput[i-1],
		})
	}
	return ll
}
