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
		ll, _ = ll.Add(i, Data{
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

	addAndCheck := func(length int, position int, data Data, t *testing.T) {
		ll := initTestLinkedList(length)
		ll, err := ll.Add(position, data)

		if err != nil {
			t.Error(err)
		}
		node, err := ll.GetNode(position)
		if node.Data.(Data).ID != data.ID {
			t.Fail()
		}
	}

	var normal = "normal "
	{
		var length = 10

		t.Run(normal+"At 1", func(t *testing.T) {
			var position = 1

			addAndCheck(length, position, testData, t)
		})

		t.Run(normal+"At middle", func(t *testing.T) {
			position := length / 2

			addAndCheck(length, position, testData, t)
		})

		t.Run(normal+"At End", func(t *testing.T) {
			var position = length + 1

			addAndCheck(length, position, testData, t)
		})
	}

	var initial = "initial"
	{
		var length = 0
		t.Run(initial+"At 1", func(t *testing.T) {
			var position = 1

			addAndCheck(length, position, testData, t)
		})
	}
}

func TestCNode_AddAtEnd(t *testing.T) {
	ll := initTestLinkedList(20)

	testData := Data{
		ID: 99,
	}

	ll = ll.AddAtEnd(testData)
	node, _ := ll.GetNode(ll.Len())
	if node.Data.(Data).ID != testData.ID {
		t.FailNow()
	}
}

func TestCNode_Delete(t *testing.T) {
	deleteAndCheck := func(length int, position int, data Data, t *testing.T) {
		ll := initTestLinkedList(length)
		node, err := ll.Delete(position)

		if err != nil {
			t.Error(err)
		}

		if node.Data.(Data).ID != data.ID {
			t.Fail()
		}
	}

	var normal = "normal "
	{
		var length = 10

		t.Run(normal+"At 1", func(t *testing.T) {
			var position = 1

			deleteAndCheck(length, position, Data{
				ID: 1,
			}, t)
		})

		t.Run(normal+"At middle", func(t *testing.T) {
			position := length / 2

			deleteAndCheck(length, position, Data{
				ID: position,
			}, t)
		})

		t.Run(normal+"At End", func(t *testing.T) {
			var position = length

			deleteAndCheck(length, position, Data{
				ID: position,
			}, t)
		})
	}

	var initial = "initial "
	{
		var length = 0
		t.Run(initial+"At 1", func(t *testing.T) {
			var position = 1

			ll := initTestLinkedList(length)
			ll, err := ll.Delete(position)

			if err != nil {
				t.Error(err)
			}

			if ll.Data != nil {
				t.Fail()
			}
		})
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
