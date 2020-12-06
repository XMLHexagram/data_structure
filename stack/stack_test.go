package stack

import (
	"testing"
)

type Data int

func TestStack(t *testing.T) {
	stack := Init()
	if !stack.IsEmpty() {
		t.Errorf("init wrong")
	}

	var testData Data = 1
	stack.Push(testData)
	if stack.IsEmpty() {
		t.Errorf("%d push wrong", testData)
	}
	if stack.Top().(Data) != 1 {
		t.Errorf("%d push wrong", testData)
	}

	testData = 2
	stack.Push(testData)
	if stack.IsEmpty() {
		t.Errorf("%d push wrong", testData)
	}
	if stack.Top().(Data) != 2 {
		t.Errorf("%d push wrong", testData)
	}

	testData = 3
	stack.Push(testData)
	if stack.IsEmpty() {
		t.Errorf("%d push wrong", testData)
	}
	if stack.Top().(Data) != 3 {
		t.Errorf("%d push wrong", testData)
	}

	if stack.Pop().(Data) != 3 {
		t.Errorf("%d pop wrong", 3)
	}
	if stack.Pop().(Data) != 2 {
		t.Errorf("%d pop wrong", 2)
	}
	if stack.Pop().(Data) != 1 {
		t.Errorf("%d pop wrong", 1)
	}

	if !stack.IsEmpty() {
		t.Errorf("pop to empty wrong")
	}
}
