package devilLanguage

import "testing"

func TestDealDevilLanguage(t *testing.T) {
	testInstance := "B(ehnxgz)B"

	testAnswer := "tsaedsaeezegexenehetsaedsae"

	resultS:=DealDevilLanguage(testInstance)

	var tempAnswer string
	for !resultS.IsEmpty() {
		tempAnswer += string(resultS.Pop().(rune))
	}

	if tempAnswer != testAnswer {
		t.Errorf("test instance wrong: %s",tempAnswer)
	}
}
