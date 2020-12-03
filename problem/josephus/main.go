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
	var code int
	ll := circularLinkedList.Init()
	fmt.Println(ll.Len())

	fmt.Print("start at:")
	_, _ = fmt.Scanf("%d", &startAt)
	fmt.Print("input:")
	for i:=1;;i++{
		_, _ = fmt.Scanf("%d", &code)
		if code == -1 {
			break
		}
		if i == 1{
			ll.Data = Data{
				ID:   i,
				Code: code,
			}
			continue
		}

		ll.AddAtEnd(Data{
			ID:   i,
			Code: code,
		})
	}

	//ll.Next.Print()
	//fmt.Println("############################")
	dealJosephus(ll, startAt)

}

func dealJosephus(linkList *circularLinkedList.CNode, startAt int) {
	var step = startAt
	var length = linkList.Len()

	//linkList.Print()

	node, _ := linkList.Delete(step)

	fmt.Println("kill:", node.Data.(Data).ID, ";code:", node.Data.(Data).Code)
	//linkList.Print()
	//node.Next.Print()
	length--

	for ; length > 0; length-- {
		step = node.Data.(Data).Code

		node, _ = node.Next.Delete(step)
		//node.Next.Print()
		//fmt.Println(node.Next.Len())
		//fmt.Println("###################")
		fmt.Println("kill:", node.Data.(Data).ID, ";code:", node.Data.(Data).Code)
	}

	//node.Next.Print()
	//data := node.Next.Next.Data
	//fmt.Println("alive:", data.(Data).ID, ";code:", data.(Data).Code)
}
