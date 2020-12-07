package josephus

import (
	"testing"
)

func TestDealJosephus(t *testing.T) {
	var testStartAt int = 20
	var testInput []int = []int{3, 1, 7, 2, 4, 8, 4}
	testAnswerList := []List{
		{
			Kill: 6,
			Code: 8,
		},
		{
			Kill: 1,
			Code: 3,
		},
		{
			Kill: 4,
			Code: 2,
		},
		{
			Kill: 7,
			Code: 4,
		},
		{
			Kill: 2,
			Code: 1,
		},
		{
			Kill: 3,
			Code: 7,
		},
		{
			Kill: 5,
			Code: 4,
		},
	}
	//testInputSlice := make([]int, 0, 100)
	//for _, v := range testInput {
	//	testInputSlice = append(testInputSlice, v)
	//}

	ll := CreateJosephus(testInput)
	list := DealJosephus(testStartAt, ll)
	for k, v := range list {
		if testAnswerList[k] != v{
			t.Errorf("answer wrong:%v",v)
		}
	}
}
