package circularLinkedList

import (
	"testing"
)

type Data struct {
	ID int
}

func initTestLinkedList(length int) *CNode {
	ll := Init()
	for i := 1; i <= length; i++ {
		_ = ll.Add(i, Data{
			ID: i,
		})
	}

	return ll
}

func TestCNode_GetNode(t *testing.T) {
	ll := initTestLinkedList(20)

	for i := 1; i <= 20; i++ {
		node, err := ll.GetNode(i)
		if err != nil {
			t.FailNow()
		}
		if node.Data.(Data).ID != i {
			t.Errorf("%d down", i)
			//t.FailNow()
		}
	}
}

func TestCNode_Add(t *testing.T) {
	testData := Data{
		ID: 99,
	}

	var length = 20

	addAndCheck := func(position int,data Data,t *testing.T) {
		ll := initTestLinkedList(length)
		err := ll.Add(position, testData)

		if err != nil {
			t.Error(err)
		}
		node ,err:= ll.GetNode(position)
		if node.Data.(Data).ID != testData.ID {
			t.Fail()
		}
		ll.Print()
	}

	var normal = "normal "
	t.Run(normal+"At 1", func(t *testing.T) {
		var position = 1

		addAndCheck(position,testData,t)
	})

	t.Run(normal+"At middle", func(t *testing.T) {
		position := length/2

		addAndCheck(position,testData,t)
	})

	t.Run(normal+"At End", func(t *testing.T) {
		var position = length +1

		addAndCheck(position,testData,t)
	})
}

func TestCNode_AddAtEnd(t *testing.T) {
	ll := initTestLinkedList(20)

	testData := Data{
		ID: 99,
	}
	ll.AddAtEnd(testData)
	node, _ := ll.GetNode(ll.Len())
	if node.Data.(Data).ID != testData.ID {
		t.FailNow()
	}
}

func TestCNode_Delete(t *testing.T) {
	ll := initTestLinkedList(20)

	deleteNode, _ := ll.Delete(ll.Len())
	if deleteNode.Data.(Data).ID != 20 {
		t.FailNow()
	}
}

func TestCNode_DeleteAtEnd(t *testing.T) {
	ll := initTestLinkedList(20)

	node := ll.DeleteAtEnd()
	if node.Data.(Data).ID != 20 {
		t.FailNow()
	}
}

func TestCNode_Len(t *testing.T) {
	ll := initTestLinkedList(20)

	if ll.Len() != 20 {
		t.FailNow()
	}
}
