package circularLinkedList

// 无头指针的循环链表

import (
	"errors"
	"fmt"
)

var (
	ErrorInvalidPosition = errors.New("invalid position")
)

type CircularLinkedList interface {
	Add(position int, data interface{}) error
	AddAtEnd(data interface{})
	Delete(position int) (*CNode, error)
	DeleteAtEnd() interface{}
	GetElem(position int) (interface{}, error)
	//todo : GetAll()
	Len() int
}

type CNode struct {
	Data interface{}
	Next *CNode
}

//func init() {
//	var a CNode
//	var b CircularLinkedList
//	b = &a
//	fmt.Println(b)
//}

func Init() *CNode {
	cNode := &CNode{
		Data: nil,
		Next: nil,
	}

	cNode.Next = cNode
	return cNode
}

func (linkList *CNode) Add(position int, data interface{}) error {
	ll := linkList

	if position < 1 {
		return ErrorInvalidPosition
	}

	if position == 1 {
		ll.Data = data
		return nil
	}

	for i := 1; i < position-1; i++ {
		ll = ll.Next
	}

	newNode := &CNode{
		Data: data,
		Next: ll.Next,
	}
	ll.Next = newNode
	return nil
}

func (linkList *CNode) Print() {
	ll := linkList

	fmt.Println("###")
	for ll.Next != linkList {
		fmt.Println(ll.Data)
		ll = ll.Next
	}

	fmt.Println(ll.Data)
	fmt.Println("###")
}

func (linkList *CNode) AddAtEnd(data interface{}) {
	_ = linkList.Add(linkList.Len()+1, data)
}

// todo： delete 1的时候的非预期行为
func (linkList *CNode) Delete(position int) (*CNode, error) {
	ll := linkList

	if position < 1 {
		return nil, ErrorInvalidPosition
	}

	//if position == 1 {
	//	position = linkList.Len() + 1
	//}

	if position == 1 {
		for ll.Next != linkList {
			ll = ll.Next
		}
	}

	for i := 1; i < position-1; i++ {
		ll = ll.Next
		//fmt.Println("*")
	}

	deleteNode := ll.Next
	ll.Next = ll.Next.Next
	return deleteNode, nil
}

func (linkList *CNode) DeleteAtEnd() interface{} {
	data, _ := linkList.Delete(linkList.Len())
	return data
}

func (linkList *CNode) GetElem(position int) (interface{}, error) {
	ll := linkList

	if position < 1 {
		return nil, ErrorInvalidPosition
	}

	for i := 0; i < position-1; i++ {
		ll = ll.Next
	}

	return ll.Data, nil
}

func (linkList *CNode) Len() int {
	ll := linkList

	length := 1
	for ll.Next != linkList {
		ll = ll.Next
		length++
	}

	return length
}
