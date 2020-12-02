package josephus

import (
	"fmt"
	"github.com/lmx-Hexagram/data_structure/link/circularLinkedList"
)

type Data struct {
	ID   int
	Code int
}

func Go() {
	// 约瑟夫环 Josephus problem
	var startAt int
	var number int
	var id int
	ll := circularLinkedList.Init()
	fmt.Println(ll.Len())

	fmt.Print("start at:")
	_, _ = fmt.Scanf("%d", &startAt)
	fmt.Print("input:")
	for {
		_, _ = fmt.Scanf("%d", &number)
		if number == -1 {
			break
		}
		id++
		ll.AddAtEnd(Data{
			ID:   id,
			Code: number,
		})

	}

	ll.Print()
	fmt.Println("############################")
	dealJosephus(ll, startAt)

}

func dealJosephus(linkList *circularLinkedList.CNode, startAt int) {
	var nowPosition = startAt
	var length = linkList.Len()

	//linkList.Print()

	node, _ := linkList.Delete(nowPosition + 1)

	fmt.Println("kill:", node.Data.(Data).ID, ";code:", node.Data.(Data).Code)
	//linkList.Print()
	length--

	for ; length > 1; length-- {
		nowPosition = node.Data.(Data).Code

		node, _ = node.Delete(nowPosition + 1)
		fmt.Println("kill:", node.Data.(Data).ID, ";code:", node.Data.(Data).Code)
	}
	data, _ := linkList.GetElem(1)
	fmt.Println("alive:", data.(Data).ID, ";code:", data.(Data).Code)
}
