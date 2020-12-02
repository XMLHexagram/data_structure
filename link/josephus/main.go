package main

import (
	"fmt"
)

type Data struct {
	ID   int
	Code int
}

type LinkNode struct {
	data     Data
	nextNode *LinkNode
}

type LinkList interface {
	initLinkList() *LinkList
	addToEnd(int) *LinkList
	print() *LinkList
	len() int
	delete(int) *LinkList
	deleteAtEnd() *LinkList
	getElem(int) int
}

func main() {
	// 约瑟夫环 Josephus problem
	linkList := initLinkList()
	var startAt int
	var number int
	id := 0

	fmt.Print("start at:")
	_, _ = fmt.Scanf("%d", &startAt)
	fmt.Print("input:")
	for {
		_, _ = fmt.Scanf("%d", &number)
		if number == -1 {
			break
		}
		id++
		linkList.addAtEnd(Data{
			ID:   id,
			Code: number,
		})
	}
	linkList.print()

	dealJosephus(linkList, startAt)
}

func dealJosephus(linkList *LinkNode, startAt int) {
	var nowPosition = startAt
	var length = linkList.len()
	fmt.Println(linkList.len())

	linkList, data := linkList.delete(nowPosition - 1)
	fmt.Println("kill:", data.ID, ";code:", data.Code)
	length--

	for ; length > 1; length-- {
		nowPosition = data.Code

		linkList, data = linkList.delete(nowPosition)
		fmt.Println("kill:", data.ID, ";code:", data.Code)
	}
	fmt.Println("alive:", linkList.getElem(1).ID, ";code:", linkList.getElem(1).Code)
}

func initLinkList() *LinkNode {
	linkList := &LinkNode{
		data: Data{
			ID:   -1,
			Code: 0,
		},
		nextNode: nil,
	}
	linkList.nextNode = linkList
	return linkList
}

func (linkList *LinkNode) addAtEnd(data Data) *LinkNode {
	linkList.add(data, linkList.len()+1)
	return linkList
}

func (linkList *LinkNode) add(data Data, position int) *LinkNode {
	if linkList.len() == 1 && linkList.nextNode == linkList && linkList.data.ID == -1 {
		linkList.data = data
		return linkList
	}
	tempLinkList := linkList

	for i := 0; i < position-1; i++ {
		if tempLinkList.nextNode == linkList {
			break
		}
		tempLinkList = tempLinkList.nextNode
	}

	linkNode := &LinkNode{
		data:     data,
		nextNode: tempLinkList.nextNode,
	}
	tempLinkList.nextNode = linkNode

	return linkList
}

func (linkList *LinkNode) getElem(position int) Data {
	tempLinkList := linkList
	for i := 0; i < position; i++ {
		if tempLinkList.nextNode == linkList {
			tempLinkList = tempLinkList.nextNode
		}
		tempLinkList = tempLinkList.nextNode
	}
	return tempLinkList.data
}

func (linkList *LinkNode) deleteAtEnd() *LinkNode {
	linkList.delete(linkList.len())
	return linkList
}

func (linkList *LinkNode) delete(position int) (*LinkNode, Data) {
	tempLinkList := linkList

	for i := 0; i < position-1; i++ {
		tempLinkList = tempLinkList.nextNode
	}
	data := tempLinkList.nextNode.data
	tempLinkList.nextNode = tempLinkList.nextNode.nextNode
	return tempLinkList, data
}

func (linkList *LinkNode) print() *LinkNode {
	tempLinkList := linkList

	for tempLinkList.nextNode != linkList {
		fmt.Println(tempLinkList.data)
		tempLinkList = tempLinkList.nextNode
	}
	fmt.Println(tempLinkList.data)
	return linkList
}

func (linkList *LinkNode) len() int {
	tempLinkList := linkList

	length := 0
	for tempLinkList.nextNode != linkList {
		length++
		tempLinkList = tempLinkList.nextNode
	}
	length++
	return length
}
