package main

import (
	"fmt"
)

type Data struct {
	ID   int
	Code int
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
