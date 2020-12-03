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
	Add(position int, data interface{}) (*CNode,error)
	AddAtEnd(data interface{}) *CNode
	Delete(position int) (*CNode, error)
	DeleteAtEnd() *CNode
	GetNode(position int) (*CNode, error)
	IsInitial() bool
	//todo : GetAll()
	Len() int
}

type CNode struct {
	Data interface{}
	Next *CNode
}

//func init() {
//	var a CircularLinkedList
//	var b CNode
//	a = &b
//	fmt.Print(a)
//}

func Init() *CNode {
	cNode := &CNode{
		Data: nil,
		Next: nil,
	}

	cNode.Next = cNode
	return cNode
}

func (linkList *CNode) IsInitial() bool {
	if linkList.Next == linkList && linkList.Data == nil {
		return true
	}
	return false
}

func (linkList *CNode) Add(position int, data interface{}) (*CNode,error) {
	ll := linkList

	if position < 1 {
		return linkList,ErrorInvalidPosition
	}

	if linkList.IsInitial() {
		ll.Data = data
		return linkList,nil
	}

	if position == 1 {
		for ll.Next != linkList {
			ll = ll.Next
		}
	}

	for i := 1; i < position-1; i++ {
		ll = ll.Next
	}

	newNode := &CNode{
		Data: data,
		Next: ll.Next,
	}
	ll.Next = newNode

	if position  == 1{
		return ll.Next,nil
	}

	return linkList,nil
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

func (linkList *CNode) AddAtEnd(data interface{}) *CNode{
	ll,_ := linkList.Add(linkList.Len()+1, data)
	return ll
}

// todo： delete 1的时候的非预期行为
func (linkList *CNode) Delete(position int) (*CNode, error) {
	ll := linkList

	if position < 1 {
		return nil, ErrorInvalidPosition
	}

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

func (linkList *CNode) DeleteAtEnd() *CNode {
	deleteNode, _ := linkList.Delete(linkList.Len())
	return deleteNode
}

func (linkList *CNode) GetNode(position int) (*CNode, error) {
	ll := linkList

	if position < 1 {
		return nil, ErrorInvalidPosition
	}

	for i := 0; i < position-1; i++ {
		ll = ll.Next
	}

	return ll, nil
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
