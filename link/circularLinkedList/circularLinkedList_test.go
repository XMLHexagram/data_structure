package main

import (
	"testing"
)

type Data struct {
	ID int
}

func TestCNode_GetElem(t *testing.T) {
	ll := InitCircularLinkedList()
	for i := 1; i <= 20; i++ {
		err := ll.Add(i, Data{
			ID: i,
		})
		if err != nil {
			t.FailNow()
		}
	}

	for i := 1; i <= 20; i++ {
		data, err := ll.GetElem(i)
		if err != nil {
			t.FailNow()
		}
		if data.(Data).ID != i {
			t.Errorf("%d down", i)
			//t.FailNow()
		}
	}
}

func TestCNode_Add(t *testing.T) {
	ll := InitCircularLinkedList()
	for i := 1; i <= 20; i++ {
		err := ll.Add(i, Data{
			ID: i,
		})
		if err != nil {
			t.FailNow()
		}
	}

	testData := Data{
		ID: 99,
	}
	_ = ll.Add(ll.Len()+1, testData)
	data, _ := ll.GetElem(ll.Len())
	if data.(Data).ID != testData.ID {
		t.FailNow()
	}
}

func TestCNode_AddAtEnd(t *testing.T) {
	ll := InitCircularLinkedList()
	for i := 1; i <= 20; i++ {
		err := ll.Add(i, Data{
			ID: i,
		})
		if err != nil {
			t.FailNow()
		}
	}

	testData := Data{
		ID: 99,
	}
	ll.AddAtEnd(testData)
	data, _ := ll.GetElem(ll.Len())
	if data.(Data).ID != testData.ID {
		t.FailNow()
	}
}

func TestCNode_Delete(t *testing.T) {
	ll := InitCircularLinkedList()
	for i := 1; i <= 20; i++ {
		err := ll.Add(i, Data{
			ID: i,
		})
		if err != nil {
			t.FailNow()
		}
	}

	data ,_ := ll.Delete(ll.Len())
	if data.(Data).ID != 20 {
		t.FailNow()
	}
}

func TestCNode_DeleteAtEnd(t *testing.T) {
	ll := InitCircularLinkedList()
	for i := 1; i <= 20; i++ {
		err := ll.Add(i, Data{
			ID: i,
		})
		if err != nil {
			t.FailNow()
		}
	}

	data := ll.DeleteAtEnd()
	if data.(Data).ID != 20 {
		t.FailNow()
	}
}

func TestCNode_Len(t *testing.T) {
	ll := InitCircularLinkedList()
	for i := 1; i <= 20; i++ {
		err := ll.Add(i, Data{
			ID: i,
		})
		if err != nil {
			t.FailNow()
		}
	}

	if ll.Len() != 20 {
		t.FailNow()
	}
}
